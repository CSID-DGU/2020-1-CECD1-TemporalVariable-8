export interface Member {
    name: string
    email: string
    address?: string
    phone?: string
    agree: {
        advertisement: boolean
        notice: boolean
        logging: boolean
    }
}
