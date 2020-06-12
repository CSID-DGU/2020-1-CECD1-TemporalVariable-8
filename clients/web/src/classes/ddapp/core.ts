export type ETHAddress = string

export interface Service {
    hint: string
    query: UniqueQuery
}

export interface DDApp extends Service{
    owner: ETHAddress
    defaultAuth : UniqueQuery
    //
    service(query: UniqueQuery): Promise<Service>
    services(query: Query): Promise<Array<Service>>
}

export abstract class Query {
    readonly typeId: string

    protected constructor(typeId: string) {
        this.typeId = typeId;
    }
    abstract clone() : Query
}

export abstract class UniqueQuery extends Query {
    abstract uclone() : UniqueQuery
}
