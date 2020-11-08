import {Service, UniqueQuery} from "@/classes/ddapp/core";
import {Option} from "fp-ts/lib/Option";


export class Authorizer {
    authorizer: UniqueQuery
    key: string

    constructor(authorizer: UniqueQuery, key: string) {
        this.authorizer = authorizer;
        this.key = key;
    }
}

export class Auth {
    authorizer: UniqueQuery
    key: string
    secret: string

    constructor(authorizer: UniqueQuery, key: string, secret: string) {
        this.authorizer = authorizer;
        this.key = key;
        this.secret = secret;
    }

    letAuthorizer(): Authorizer {
        return new Authorizer(this.authorizer.uclone(), this.key)
    }
}

export interface AuthService extends Service {
    login(id: string, pw: string): Promise<Option<Auth>>;

    authorize(a: Authorizer, t: Auth): Promise<{ isAuthorize: boolean; msg?: string }>;
}
