import Vue from 'vue';

export default interface IConfirm extends Vue {
    open(handle?: (arg0: boolean) => void): void
}