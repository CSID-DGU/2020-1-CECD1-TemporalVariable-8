import {DDApp, ETHAddress, Query, Service, UniqueQuery} from "@/classes/ddapp/core";
import {Auth, Authorizer, AuthService} from "@/classes/ddapp/authorize";
import {DDAppQuery, IDQuery, InQuery} from "@/classes/ddapp/query";
import {v4} from "uuid"
import * as opt from "fp-ts/lib/Option";
import {Menu, Order, OrderID, OrderInfo, Progress, RestaurantService} from "@/classes/ddapp/restaurant";
import {Either, left, right} from "fp-ts/lib/Either";
import {NonEmptyArray} from "fp-ts/lib/NonEmptyArray";
import {isNonEmpty} from "fp-ts/lib/Array";
import assert from "assert";
import {isSome, none, Option, some} from "fp-ts/lib/Option";

export class MockAccount {
    constructor(id: string, pw: string) {
        this.id = id;
        this.pw = pw;
    }

    id: string
    pw: string
}

export const USERS = {
    OWNER: new MockAccount("owner", "PW_OWNER"),
    SELLER_0: new MockAccount("seller0", "PW_SELLER_0"),
    SELLER_1: new MockAccount("seller1", "PW_SELLER_1"),
    SELLER_2: new MockAccount("seller2", "PW_SELLER_2"),
    SELLER_3: new MockAccount("seller3", "PW_SELLER_3"),
    CUNSUMER: new MockAccount("consumer", "PW_CUNSUMER"),
}
export const SERVICES = {
    DDAPP: {
        ethAddress: "MOCK, ddApp",
        hint: "hint : MOCK, ddApp",
    },
    AUTH: {
        ethAddress: "MOCK, AUTH",
        hint: "hint : MOCK, AUTH"
    },
    SELLER_CHINAFOOD: {
        ethAddress: "MOCK, SELLER_CHINAFOOD",
        hint: "hint : MOCK, SELLER_CHINAFOOD",
        name: "중국집",
        menus: [
            {name: "짜장면", desc: "새까만 짜장면!"}
        ]
    },
    SELLER_ITALIFOOD: {
        ethAddress: "MOCK, SELLER_ITALIFOOD",
        hint: "hint : MOCK, SELLER_ITALIFOOD",
        name: "이탈리아집"
    },
    GENERATOR: (): string => {
        return v4()
    }
}

class MockOrder implements Order {
    count: number;
    coupon?: string;

    constructor() {
        this.count = 1;
    }

}

class MockProgress implements Progress {
    msg?: string;
    percent?: number;

    constructor(msg?: string, percent?: number) {
        this.msg = msg;
        this.percent = percent;
    }

}

class MockMenu extends Menu {
    constructor(name: string, desc: string, image?: URL) {
        super(v4(), name, desc, image);
    }

    buildOrder(props: OrderInfo): Either<NonEmptyArray<string>, Order> {
        const errs: Array<string> = []
        if (props.count !== undefined) {
            if (props.count < 0) {
                errs.push("count must have at least 1")
            }
            if (!Number.isInteger(props.count)) {
                errs.push("count must be integer")
            }
        }
        if (isNonEmpty(errs)) {
            return left(errs)
        }
        //
        const temp = new MockOrder()
        if (props.count !== undefined) {
            temp.count = props.count
        }
        temp.coupon = props.coupon
        return right(temp)
    }

}

export class MockRestaurant implements RestaurantService {
    hint: string;
    name: string;
    owner: Authorizer;
    query: UniqueQuery;
    _menus: Array<Menu> = []
    _orders: Map<OrderID, [Order, Progress]> = new Map()

    constructor(hint: string, name: string, owner: Authorizer, query: UniqueQuery) {
        this.hint = hint;
        this.name = name;
        this.owner = owner;
        this.query = query;
    }


    menus(): Promise<Array<Menu>> {
        return Promise.resolve(this._menus);
    }

    progression(p: OrderID): Promise<Progress> {
        const found = this._orders.get(p)
        if (found !== undefined) {
            return Promise.resolve(found[1]);
        }
        return Promise.reject(`Not Found ${p}`)
    }

    request(request: Order): Promise<OrderID> {
        const oid = new OrderID(this.query, v4())
        // 여기서 원래 주문이 가능한지 검증하여야 함
        //
        this._orders.set(oid, [request, new MockProgress("start", 0.0)])
        return Promise.resolve(oid);
    }

}

class MockAuthService implements AuthService {
    hint = SERVICES.AUTH.hint
    query: UniqueQuery = new IDQuery(SERVICES.AUTH.ethAddress)
    users: Array<MockAccount>

    constructor(users: Array<MockAccount>) {
        this.users = users
    }

    authorize(a: Authorizer, t: Auth): Promise<{ isAuthorize: boolean; msg?: string }> {
        return Promise.resolve({isAuthorize: false});
    }

    login(id: string, pw: string): Promise<Option<Auth>> {
        for (const usersKey in USERS) {
            if (id === USERS[usersKey].id && pw === USERS[usersKey].pw) {
                return Promise.resolve(some(new Auth(this.query.uclone(), id, pw)))
            }
        }
        return Promise.resolve(none);
    }

}

export class MockDDApp implements DDApp {
    owner: ETHAddress = SERVICES.DDAPP.ethAddress
    hint = SERVICES.DDAPP.hint
    query: UniqueQuery = new DDAppQuery()
    defaultAuth: UniqueQuery = new IDQuery(SERVICES.AUTH.ethAddress);
    _defaultAuthSvr: AuthService
    _services: Map<ETHAddress, Service>

    constructor() {
        this._defaultAuthSvr = new MockAuthService(Array.from((Object.values(USERS)).entries()).map(value => value[1]))
        this._services = new Map<ETHAddress, Service>([
            [SERVICES.DDAPP.ethAddress, this],
            [SERVICES.AUTH.ethAddress, this._defaultAuthSvr],
        ])
    }

    async init() {
        try {
            const s0 = await this._defaultAuthSvr.login(USERS.SELLER_0.id, USERS.SELLER_0.pw)
            if (!isSome(s0)) {
                return
            }
            // 중국집 추가
            const restC = new MockRestaurant(
                SERVICES.SELLER_CHINAFOOD.hint,
                SERVICES.SELLER_CHINAFOOD.name,
                s0.value.letAuthorizer(),
                new IDQuery(SERVICES.SELLER_CHINAFOOD.ethAddress)
            )
            for (const menu of SERVICES.SELLER_CHINAFOOD.menus) {
                restC._menus.push(new MockMenu(menu.name, menu.desc))
            }
            this._services.set(SERVICES.SELLER_CHINAFOOD.ethAddress, restC)
        } catch (e) {
            assert.fail(e)
        }
    }

    service(query: UniqueQuery): Promise<Service> {
        if (query instanceof DDAppQuery) {
            return Promise.resolve(this);
        }
        if (query instanceof IDQuery) {
            return opt.fold<Service, Promise<Service>>(
                () => Promise.reject<Service>(`Not Found ID '${query.id}'`),
                a => Promise.resolve(a)
            )(opt.fromNullable(this._services.get(query.id)));
        }
        return Promise.reject(`Not Found : ${query}`);
    }

    services(query: Query): Promise<Array<Service>> {
        const res = MockDDApp.parseQuery(Array.from(this._services.entries()), query)
        if (typeof res === 'string') {
            return Promise.reject(res)
        } else {
            return Promise.resolve(res.map(value => value[1]))
        }
    }

    private static parseQuery(gh: Array<[ETHAddress, Service]>, q: Query): Array<[ETHAddress, Service]> | string {
        switch (q.typeId) {
            case DDAppQuery.name: {
                return gh.filter(value => value[1] instanceof MockDDApp)
            }
            case IDQuery.name: {
                return gh.filter(value => value[0] === (q as IDQuery).id)
            }
            case InQuery.name: {
                const iq = (q as InQuery)
                const vs = iq.querys.map(value => this.parseQuery(gh, value))
                const err = vs.find(value => typeof value === 'string')
                if (err !== undefined) {
                    return err
                }
                return vs.flatMap(value => value as Array<[ETHAddress, Service]>)
            }
        }
        return `Ununknown Query '${q}'`
    }

}


// export class MockOrderProgress implements OrderProgression {
//     constructor(order: Order) {
//         this.order = order;
//     }
//
//     order: Order;
//
// }
//
// export class MockRestaurant extends Restaurant {
//     owner: Auth<undefined>;
//     identifier: Query;
//     name: string;
//     hint: string;
//     address?: string;
//     private _menus: Array<Menu>
//
//     constructor(owner: Auth<undefined>, identifier: Query, name: string, hint: string) {
//         super();
//         this.identifier = identifier
//         this.owner = owner;
//         this.name = name;
//         this.hint = hint;
//         this._menus = [];
//     }
//
//     addMenu(m: Menu): Promise<Menu> {
//         return new Promise<Menu>((resolve, reject) => {
//             resolve(m)
//         })
//     }
//
//     menus(): Array<Menu> {
//         return this._menus;
//     }
//
//     request(order: Order): Promise<OrderProgression> {
//         return new Promise<OrderProgression>(
//             (resolve, reject) => {
//                 resolve(new MockOrderProgress(order))
//             }
//         );
//     }
//
//
// }
//
// function createMRestaurant(a: MockRestaurant, f: (mr: MockRestaurant) => void): MockRestaurant {
//     f(a)
//     return a
// }
//
// export class MockAuth<T> implements Auth<T> {
//
//     constructor(authServiceQuery: Query, identifier: string, isAuthorized: boolean, isOwner: boolean, secret?: T) {
//         this.authServiceQuery = authServiceQuery;
//         this.isAuthorized = isAuthorized;
//         this.isOwner = isOwner;
//         this.identifier = identifier;
//         this.secret = secret
//     }
//     authServiceQuery: Query;
//     isAuthorized: boolean;
//     isOwner: boolean;
//     identifier: string;
//     secret?: T;
// }
//
//
// export class MockBlockchain implements Blockchain, Service {
//     address: EthAddress;
//     hint: string;
//     identifier: Query;
//     owner: Auth<undefined>;
//     _services: Array<Service>;
//     _account: Array<{ id: string; pw: string; isOwner: boolean; isAuthorized: boolean }>;
//
//     constructor() {
//         //
//         this.address = "MOCK BLOCKCHAIN";
//         this.hint = "Blockchain Interface for Mock"
//         this.identifier = new BlockchainQuery(this.address)
//         this.owner = new MockAuth<undefined>(
//             this.identifier,
//             OWNER.id,
//             true,
//             true
//         )
//         this._account = [
//             OWNER,
//             SELLER_0,
//             SELLER_1,
//             SELLER_2,
//             SELLER_3,
//             CUNSUMER,
//         ]
//         this._services = [
//             this,
//             createMRestaurant(
//                 new MockRestaurant(, "한국집", "두유노 킴취?"),
//                 mr => {
//                     return null
//                 }
//             ),
//             createMRestaurant(
//                 new MockRestaurant({
//                     ownerId: SELLER_0.id,
//                     authServiceQuery: this.identifier
//                 }, "중국집", "사실 우리는 군만두만 팔아요!"),
//                 mr => {
//
//                     return null
//                 }
//             ),
//             createMRestaurant(
//                 new MockRestaurant({ownerId: SELLER_0.id, authServiceQuery: this.identifier}, "일본집", "스시다해, 스시"),
//                 mr => {
//
//                     return null
//                 }
//             ),
//             createMRestaurant(
//                 new MockRestaurant({ownerId: SELLER_0.id, authServiceQuery: this.identifier}, "이탈리아집", "피자는 역시 파인애플"),
//                 mr => {
//
//                     return null
//                 }
//             ),
//             createMRestaurant(
//                 new MockRestaurant({ownerId: SELLER_0.id, authServiceQuery: this.identifier}, "미국집", "살찌는 음식 팝니다"),
//                 mr => {
//
//                     return null
//                 }
//             ),
//         ]
//     }
//
//     login(id: string, pw: string): Promise<Auth<string>> {
//         return new Promise((resolve, reject) => {
//             for (const accountElement of this._account) {
//                 if (accountElement.id === id && accountElement.pw === pw) {
//                     return resolve(new MockAuth(this.identifier, id, accountElement.isAuthorized, accountElement.isOwner, pw))
//                 }
//             }
//             return resolve(new MockAuth(this.identifier, id, false, false));
//         })
//     }
//
//     services(...querys: Query[]): Promise<Array<Service>> {
//         return new Promise<Array<Service>>((resolve, reject) => {
//             const res = this._services.filter(value => MockBlockchain.checkServiceIsQueryFulfilled(value, querys))
//             resolve(res)
//         });
//     }
//
//     static checkServiceIsQueryFulfilled(s: Service, querys: Query[]): boolean {
//         for (const query of querys) {
//             if (query instanceof BlockchainQuery) {
//                 if (!(s.identifier instanceof BlockchainQuery && s.identifier.target === query.target)) {
//                     return false
//                 }
//             } else if (query instanceof IdentQuery) {
//                 if (s.owner.authServiceQuery !== query) {
//                     return false
//                 }
//             } else if (query instanceof GeometryQuery) {
//                 // TODO : Geo Query Cond
//                 return false
//             }
//             // Unknown Query Type
//         }
//         return true
//     }
//
//     authorize<T>(auth: Auth<T>): Promise<Option<Auth<T>>> {
//         return Promise.resolve(none);
//     }
//
// }