import {Service, UniqueQuery} from "@/classes/ddapp/core";
import {Authorizer} from "@/classes/ddapp/authorize";
import {Either} from "fp-ts/lib/Either";
import {fold, fromNullable} from "fp-ts/lib/Option";
import {NonEmptyArray} from "fp-ts/lib/NonEmptyArray";

export abstract class RestaurantService implements Service {
    hint: string
    name: string
    query: UniqueQuery
    owner: Authorizer
    //
    avatar?: URL
    address?: string
    description?: string
    phone?: string
    link?: URL
    email?: string

    protected constructor(owner: Authorizer, query: UniqueQuery, hint: string, name: string) {
        this.owner = owner
        this.hint = hint
        this.name = name
        this.query = query
    }

    abstract menus(): Promise<Array<Menu>>

    abstract request(request: Order): Promise<OrderID>

    abstract progression(p: OrderID): Promise<Progress>
}


export abstract class Menu {
    uid: string
    name: string
    desc: string
    image?: URL

    protected constructor(uid: string, name: string, desc: string, image?: URL) {
        this.uid = uid;
        this.name = name;
        this.desc = desc;
        this.image = image;
    }

    abstract buildOrder(props: OrderInfo): Either<NonEmptyArray<string>, Order>
}

export abstract class Order {
    count = 1
    coupon?: string
}

export interface OrderInfo extends OrderInfoCount, OrderInfoCoupon {
}

export interface OrderInfoCount {
    count?: number;
}

export interface OrderInfoCoupon {
    coupon?: string;
}

export class OrderID {
    constructor(restaurant: UniqueQuery, localId: string) {
        this.restaurant = restaurant;
        this.localId = localId;
    }

    restaurant: UniqueQuery
    localId: string
}

export interface Progress {
    percent?: number;
    msg?: string;
}