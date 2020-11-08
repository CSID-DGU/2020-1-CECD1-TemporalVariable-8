export default {
    email: (testee: string): boolean | string => {
        // https://www.regular-expressions.info/email.html
        // 매우 특이한 이메일, 혹은 표준 라틴 문자로 이뤄지지 않은 이메일 주소의 경우 에러가 남(도메인이건 이름이건)
        return /^[a-z0-9!#$%&'*+/=?^_‘{|}\-~]+(?:\.[a-z0-9!#$%&'*+/=?^_‘{|}\-~]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/i.test(testee)
    }
}