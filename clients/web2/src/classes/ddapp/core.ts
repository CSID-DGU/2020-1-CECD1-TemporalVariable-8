import {ServiceType} from "@/classes/ddapp/query";

export type ETHAddress = string

export interface Service {
    hint: string;
    query: UniqueQuery;

    isService(s: ServiceType): boolean;
}

export interface DDApp extends Service {
    owner: ETHAddress;
    defaultAuth: UniqueQuery;

    //
    service(query: UniqueQuery): Promise<Service>;

    services(query: Query): Promise<Array<Service>>;
}

export abstract class Query {
    readonly typeId: string;

    protected constructor(typeId: string) {
        this.typeId = typeId;
    }

    abstract clone(): Query
}

export abstract class UniqueQuery extends Query {
    abstract uclone(): UniqueQuery

    abstract key(): string
}
