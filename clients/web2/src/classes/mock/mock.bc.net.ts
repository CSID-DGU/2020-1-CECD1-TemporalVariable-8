import {DDApp, ETHAddress, Query, Service, UniqueQuery} from "@/classes/ddapp/core";
import {Auth, Authorizer, AuthService} from "@/classes/ddapp/authorize";
import {AllQuery, DDAppQuery, IDQuery, InQuery, ServiceType, TypeQuery} from "@/classes/ddapp/query";
import {v4} from "uuid"
import * as opt from "fp-ts/lib/Option";
import {Menu, Order, OrderID, OrderInfo, Progress, RestaurantService} from "@/classes/ddapp/restaurant";
import {Either, left, right} from "fp-ts/lib/Either";
import {NonEmptyArray} from "fp-ts/lib/NonEmptyArray";
import {isNonEmpty} from "fp-ts/lib/Array";
import assert from "assert";
import {isSome, none, Option, some} from "fp-ts/lib/Option";

// export const USERS = {
//     OWNER: new Auth(, "owner", "PW_OWNER"),
//     SELLER_0: new Auth("seller0", "PW_SELLER_0"),
//     SELLER_1: new Auth("seller1", "PW_SELLER_1"),
//     SELLER_2: new Auth("seller2", "PW_SELLER_2"),
//     SELLER_3: new Auth("seller3", "PW_SELLER_3"),
//     CUNSUMER: new Auth("consumer", "PW_CUNSUMER"),
// }
// export const SERVICES = {
//     DDAPP: {
//         ethAddress: "MOCK, ddApp",
//         hint: "hint : MOCK, ddApp",
//     },
//     AUTH: {
//         ethAddress: "MOCK, AUTH",
//         hint: "hint : MOCK, AUTH"
//     },
//     SHOPS:[
//         new MockRestaurant(
//             "hint : MOCK, SELLER_CHINAFOOD",
//             "중국집",
//             USERS.SELLER_0.
//         ),
//         {
//             ethAddress: "MOCK, SELLER_CHINAFOOD",
//             hint: "hint : MOCK, SELLER_CHINAFOOD",
//             name: "중국집",
//             menus: [
//                 {name: "짜장면", desc: "새까만 짜장면!"}
//             ]
//         },
//         {
//             ethAddress: "MOCK, SELLER_ITALIFOOD",
//             hint: "hint : MOCK, SELLER_ITALIFOOD",
//             name: "이탈리아집",
//             menus: [
//                 {name: "스파게티", desc: "비싸"}
//             ]
//         },
//     ],
// }

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

    isService(s: ServiceType) {
        switch (s) {
            case ServiceType.Restaurant:
                return true
            default:
                return false
        }
    }
}

class MockAuthService implements AuthService {
    hint = "목 객체 인증 서비스"
    query: UniqueQuery = new IDQuery("MockAuthService, ETHAddress")
    users: Array<{ id: string; pw: string }>

    constructor(users?: Array<{ id: string; pw: string }>) {
        this.users = users !== undefined ? users : []
    }

    authorize(a: Authorizer, t: Auth): Promise<{ isAuthorize: boolean; msg?: string }> {
        return Promise.resolve({isAuthorize: false});
    }

    login(id: string, pw: string): Promise<Option<Auth>> {
        for (const usersKey of this.users) {
            if (id === usersKey.id && pw === usersKey.pw) {
                return Promise.resolve(some(new Auth(this.query.uclone(), id, pw)))
            }
        }
        return Promise.resolve(none);
    }

    isService(s: ServiceType) {
        switch (s) {
            case ServiceType.Auth:
                return true
            default:
                return false
        }
    }
}

export class MockDDApp implements DDApp {
    owner: ETHAddress = "MockDDapp, ETHAddress"
    hint = "목 객체 ddApp"
    query: UniqueQuery = new DDAppQuery()
    defaultAuth: UniqueQuery = new IDQuery("MockAuthorizer, ETHAddress");
    _defaultAuthSvr: MockAuthService
    _services: Map<ETHAddress, Service>

    constructor() {
        this._defaultAuthSvr = new MockAuthService([])
        this._services = new Map<ETHAddress, Service>([
            [this.owner, this],
            [(this._defaultAuthSvr.query as IDQuery).id, this._defaultAuthSvr],
        ])
    }

    //
    addUser(id: string, pw: string): Either<string, Auth>{
        if (this._defaultAuthSvr.users.every(value => value.id !== id)){
            this._defaultAuthSvr.users.push({id: id, pw: pw})
            return right(new Auth(this._defaultAuthSvr.query.uclone(), id, pw))
        }
        return left("존재하는 ID")
    }
    addRestaurant(name: string, hint: string, owner: string | Auth, ethID: ETHAddress): Either<string, MockRestaurant>{
        if (this._services.has(ethID)){
            return left("이더리움 어드레스 충돌")
        }
        const account = this._defaultAuthSvr.users.find(value => value.id === owner)
        if (account === undefined){
            return left("존재하지 않는 소유주")
        }
        const res = new MockRestaurant(
            hint,
            name,
            typeof owner === 'string' ? new Authorizer(this._defaultAuthSvr.query.uclone(), owner) : owner.letAuthorizer(),
            new IDQuery(ethID),
        )
        this._services.set(ethID, res)
        return right(res)

    }

    //

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
            case AllQuery.name: {
                return gh
            }
            case TypeQuery.name: {
                const esvr = (q as TypeQuery).st
                return gh.filter(value => value[1].isService(esvr))
            }
        }
        return `Ununknown Query '${q}'`
    }

    isService(s: ServiceType) {
        switch (s) {
            case ServiceType.DDApp:
            case ServiceType.Auth:
                return true
            default:
                return false
        }
    }
}