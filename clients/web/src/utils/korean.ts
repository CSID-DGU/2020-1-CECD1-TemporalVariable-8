import {inRange} from 'lodash'

const FIRST_HANGUL = 0xAC00
const LAST_HANGUL = 0xD7A3

// 종성을 가진지 판단
function hasJongseong(cc: number): boolean {
    return (cc - FIRST_HANGUL) % 28 > 0
}

// 한글인지 확인
function isHangul(cc: number): boolean {
    return inRange(cc, FIRST_HANGUL, LAST_HANGUL)
}

export const investigation = {
    subject: (subjectword: string) => {
        const lastchar = subjectword.charCodeAt(subjectword.length - 1)
        if (isHangul(lastchar)) {
            if (hasJongseong(lastchar)) {
                return '이'
            } else {
                return '가'
            }
        }
        return ''
    },
    objective: (objectiveword: string) => {
        const lastchar = objectiveword.charCodeAt(objectiveword.length - 1)
        if (isHangul(lastchar)) {
            if (hasJongseong(lastchar)) {
                return '을'
            } else {
                return '를'
            }
        }
        return ''
    },
}

