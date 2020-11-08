import {ETHAddress, Query, Service, UniqueQuery} from "@/classes/ddapp/core";

export class DDAppQuery implements UniqueQuery {
    readonly typeId: string = DDAppQuery.name;

    clone(): Query {
        return this;
    }

    uclone(): UniqueQuery {
        return this;
    }

    key(): string {
        return `${DDAppQuery.name}`;
    }
}

export class IDQuery implements UniqueQuery {
    readonly typeId: string = IDQuery.name;
    readonly id: ETHAddress;

    //
    constructor(id: ETHAddress) {
        this.id = id;
    }

    //
    clone(): Query {
        return this;
    }

    uclone(): UniqueQuery {
        return this;
    }

    key(): string {
        return `${UniqueQuery.name}-${this.id}`;
    }
}

export class InQuery implements Query {
    readonly typeId: string = InQuery.name
    querys: Array<Query>

    constructor(...querys: Array<Query>) {
        this.querys = querys;
    }

    clone(): Query {
        return new InQuery(...this.querys.map(value => value.clone()));
    }
}

export class AllQuery implements Query {
    readonly typeId: string = AllQuery.name

    clone(): Query {
        return this
    }
}

export class TypeQuery implements Query {
    readonly typeId: string = TypeQuery.name
    st: ServiceType

    clone(): Query {
        return this
    }

    constructor(c: ServiceType) {
        this.st = c;
    }
}

export enum ServiceType {
    DDApp,
    Auth,
    Restaurant,
}