import {Auth, Authorizer, AuthService} from "@/classes/ddapp/authorize";
import {UniqueQuery} from "@/classes/ddapp/core";
import {Option} from "fp-ts/lib/Option";
import {ServiceType} from "@/classes/ddapp/query";

export class NetCtlAuthService<T extends AuthService> implements AuthService {
    hint: string
    query: UniqueQuery
    t: T
    msg?: string

    constructor(t: T) {
        this.hint = t.hint;
        this.query = t.query;
        this.t = t;
    }

    setError(msg: string) {
        this.msg = msg
    }

    resetError() {
        this.msg = undefined
    }

    authorize(a: Authorizer, t: Auth): Promise<{ isAuthorize: boolean; msg?: string }> {
        if (this.msg === undefined) {
            return this.t.authorize(a, t)
        }
        return Promise.reject(this.msg)
    }

    login(id: string, pw: string): Promise<Option<Auth>> {
        if (this.msg === undefined) {
            return this.t.login(id, pw)
        }
        return Promise.reject(this.msg)
    }

    isService(s: ServiceType): boolean {
        return this.t.isService(s);
    }

}