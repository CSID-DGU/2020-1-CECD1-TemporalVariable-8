import {ETHAddress, Query, UniqueQuery} from "@/classes/ddapp/core";

export class DDAppQuery implements UniqueQuery{
    readonly typeId: string = IDQuery.name
    constructor() {
    }
    clone(): Query {
        return this;
    }
    uclone(): UniqueQuery {
        return this;
    }
}
export class IDQuery implements UniqueQuery{
    readonly typeId: string = IDQuery.name
    readonly id : ETHAddress
    constructor(id: ETHAddress) {
        this.id = id;
    }
    clone(): Query {
        return this;
    }
    uclone(): UniqueQuery {
        return this;
    }
}
export class InQuery implements Query{
    readonly typeId: string = InQuery.name
    querys : Array<Query>
    constructor(...querys: Array<Query>) {
        this.querys = querys;
    }
    clone(): Query {
        return new InQuery(...this.querys.map(value => value.clone()));
    }
}