import {MockDDApp, SERVICES, USERS} from "@/classes/mock/mock.bc.net";
import {DDAppQuery} from "@/classes/ddapp/query";
import chai, {expect} from "chai"
import chaiAsPromised from "chai-as-promised"
import {AuthService} from "@/classes/ddapp/authorize";
import {isSome, none} from "fp-ts/lib/Option";
import {NetCtlAuthService} from "@/classes/mock/utils";

chai.use(chaiAsPromised)

describe('BlockChain Mock', () => {
    const MBC = new MockDDApp()
    before("목 객체 초기화 동작, 각종 서비스 추가 등", async () => {
        return await MBC.init()
    })
    describe("쿼리 테스트", ()=>{

        it('쿼리 : ddApp 단일 검색', function () {
            return expect(MBC.service(new DDAppQuery())).eventually.fulfilled.eq(MBC)
        });
        it('쿼리 : 표준 인증장치 단일 검색', function () {
            return expect(MBC.service(MBC.defaultAuth))
                .eventually.fulfilled.include({
                    hint : SERVICES.AUTH.hint,
                })
        });
    })
    describe("로그인 테스트", () => {
        it('로그인 : 네트워크 지연 생성 ', () => {
            return expect(
                MBC.service(MBC.defaultAuth)
                    .then(value => value as AuthService)
                    .then(value => {
                        const tmp = new NetCtlAuthService(value)
                        tmp.setError("사용자 네트워크 오류")
                        return tmp
                    })
                    .then(value => value.login("INVALID", "INVALID"))
            ).eventually.rejectedWith("사용자 네트워크 오류")
        });
        it('로그인 : 표준 인증장치, 잘못된 계정으로 ', () => {
            return expect(
                MBC.service(MBC.defaultAuth)
                    .then(value => value as AuthService)
                    .then(value => value.login("INVALID", "INVALID"))
            ).eventually.fulfilled.eq(none)
        });
        it('로그인 : 표준 인증장치, 올바른 계정, 잘못된 비밀번호 ', () => {
            return expect(
                MBC.service(MBC.defaultAuth)
                    .then(value => value as AuthService)
                    .then(value => value.login(USERS.CUNSUMER.id, "INVALID"))
            ).eventually.fulfilled.eq(none)
        });
        it('로그인 : 표준 인증장치, 올바른 계정, 올바른 비밀번호 ', () => {
            return expect(
                MBC.service(MBC.defaultAuth)
                    .then(value => value as AuthService)
                    .then(value => value.login(USERS.CUNSUMER.id, USERS.CUNSUMER.pw))
            ).eventually.fulfilled
                .satisfy(isSome)
                .property("value").include({
                    key: USERS.CUNSUMER.id,
                    secret: USERS.CUNSUMER.pw,
                })
        });
    })
})

