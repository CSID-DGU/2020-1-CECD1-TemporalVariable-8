<template>
    <v-container class="pt-0">
        <v-row>
            <v-img :aspect-ratio="8/3" max-height="420" :src="require('@/assets/MakItGo.svg')"></v-img>
        </v-row>
        <v-row>
            <v-col cols="12" md="12">
                <v-window v-model='step' @change="windowChangedHandle">
                    <!-- 약관 -->
                    <v-window-item :value="0">
                        <v-form ref="form0" v-model="formModel[0]">
                            <v-container>
                                <v-row>
                                    <v-card width="100%" flat>
                                        <v-card-title>{{$t('dialog.register.termOfUse')}}</v-card-title>
                                        <v-divider class="mb-2"/>
                                        <v-card-text class="grey lighten-2 rounded"
                                                     style="white-space: pre-line ; height: 150px; overflow-y: auto">
                                            {{$t('dialog.register.termOfUseContent')}}
                                        </v-card-text>
                                    </v-card>
                                </v-row>
                                <v-row>
                                    <v-spacer/>
                                    <v-checkbox
                                            on-icon="mdi-check-circle" off-icon="mdi-check-circle-outline"
                                            color="secondary"
                                            :label="$t('dialog.register.termOfUseAgree')"
                                            :rules="[v=>!!v||$t('dialog.register.requireAgree')]"
                                            v-model="agreeTermOfUse"/>
                                </v-row>
                                <v-row>
                                    <v-card width="100%" flat>
                                        <v-card-title>{{$t('dialog.register.personalData')}}</v-card-title>
                                        <v-divider class="mb-2"/>
                                        <v-card-text class="grey lighten-2 rounded"
                                                     style="white-space: pre-line ; height: 150px; overflow-y: auto">
                                            {{$t('dialog.register.personalDataContent')}}
                                        </v-card-text>
                                    </v-card>
                                </v-row>
                                <v-row>
                                    <v-spacer/>
                                    <v-checkbox
                                            on-icon="mdi-check-circle" off-icon="mdi-check-circle-outline"
                                            color="secondary"
                                            :label="$t('dialog.register.personalDataAgree')"
                                            :rules="[v=>!!v||$t('dialog.register.requireAgree')]"
                                            v-model="agreePersonalData"/>

                                </v-row>
                            </v-container>
                        </v-form>
                    </v-window-item>
                    <!-- 기본정보 -->
                    <v-window-item :value="1">
                        <!-- https://developers.google.com/web/updates/2015/06/checkout-faster-with-autofill -->
                        <!-- 자동완성 지원을 위해 이를 준비 -->
                        <v-form ref="form1" v-model="formModel[1]">
                            <v-text-field
                                    name="name"
                                    autocomplete="name"

                                    v-model="name"
                                    :rules="ruleName"
                                    :label="$t('word.name')"
                            />
                            <v-autocomplete
                                    name="email"
                                    autocomplete="email"

                                    v-model="email"
                                    hide-no-data hide-selected
                                    clearable

                                    :rules="ruleEmail"
                                    :items="availableEmails"
                                    :search-input.sync="emailSearching"
                                    :label="$t('word.email')"
                            />
                            <v-text-field
                                    name="phone"
                                    autocomplete="tel"

                                    v-model="phone"
                                    :label="$t('word.phone')"
                            />
                            <v-text-field
                                    name="address"
                                    autocomplete="address"

                                    v-model="address"
                                    :label="$t('word.address')"
                            />
                        </v-form>
                    </v-window-item>
                    <!-- 확장정보 -->
                    <v-window-item :value="2">
                        <v-form ref="form2" v-model="formModel[2]">
                            <v-container style="min-height: 600px">
                                <v-row class="">
                                    <v-col cols="12" md="4">
                                        <video class="what-is-video" autoplay loop>
                                            <source src="/placeholder0.local.mp4" type="video/mp4">
                                        </video>
                                    </v-col>
                                    <v-col cols="12" md="8" class="d-flex flex-column">
                                        <p class="mb-auto what-is-text">
                                            {{$t('dialog.register.whatIsNoticeAllow')}}</p>
                                        <v-checkbox v-model="acceptNotice" class="align-self-end"
                                                    :label="$t('dialog.register.noticeAllow')"/>
                                    </v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="12" md="4">
                                        <video class="what-is-video" autoplay loop>
                                            <source src="/placeholder1.local.mp4" type="video/mp4">
                                        </video>
                                    </v-col>
                                    <v-col cols="12" md="8" class="d-flex flex-column">
                                        <p class="mb-auto what-is-text">
                                            {{$t('dialog.register.whatIsADAllow')}}</p>
                                        <v-checkbox v-model="acceptAD" class="align-self-end"
                                                    :label="$t('dialog.register.ADAllow')"/>
                                    </v-col>
                                </v-row>
                                <v-row>
                                    <v-col cols="12" md="4">
                                        <video class="what-is-video" autoplay loop>
                                            <source src="/placeholder2.local.mp4" type="video/mp4">
                                        </video>
                                    </v-col>
                                    <v-col cols="12" md="8" class="d-flex flex-column">
                                        <p class="mb-auto what-is-text">
                                            {{$t('dialog.register.whatIsDebugLogAllow')}}</p>
                                        <v-checkbox v-model="acceptDLog" class="align-self-end"
                                                    :label="$t('dialog.register.debugLogAllow')"/>
                                    </v-col>
                                </v-row>
                            </v-container>
                        </v-form>
                    </v-window-item>
                    <!-- 완료 -->
                    <v-window-item :value="3">
                        <v-form ref="form3" v-model="formModel[3]">
                            <h3 class="text-center">{{$t('dialog.register.registerAlmostComplete')}}</h3>
                            <v-checkbox label="임시"/>
                        </v-form>
                    </v-window-item>
                </v-window>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="1" md="2">
                <v-btn class="float-left justify-start"
                       text
                       :color="prevColor"
                       @click="prevStep">
                    <v-icon class="mr-1">{{prevIcon}}</v-icon>
                    <v-scroll-x-reverse-transition>
                        <span v-show="$vuetify.breakpoint.mdAndUp && prevMessage.length > 0">{{prevMessage}}</span>
                    </v-scroll-x-reverse-transition>
                </v-btn>
            </v-col>
            <v-col cols="10" md="8">
                <v-item-group
                        v-model="step"
                        class="text-center"
                        mandatory
                >
                    <v-item
                            v-for="n in stepMax"
                            :key="`btn-${n-1}`"
                            v-slot:default="{ active, toggle }"
                    >
                        <v-btn
                                :input-value="active"
                                icon
                                depressed
                                :disabled="!jumpActive[n-1]"
                                @click="toggle"
                        >
                            <v-icon>mdi-record</v-icon>
                        </v-btn>
                    </v-item>
                </v-item-group>
            </v-col>
            <v-col cols="1" md="2">
                <v-btn class="float-right justify-end"
                       text
                       :color="nextColor"
                       @click="nextStep">
                    <v-scroll-x-transition>
                        <span v-show="$vuetify.breakpoint.mdAndUp && nextMessage.length > 0">{{nextMessage}}</span>
                    </v-scroll-x-transition>
                    <v-icon class="ml-1">{{nextIcon}}</v-icon>
                </v-btn>
            </v-col>
        </v-row>
        <Confirm
                ref="cancelDialog"
                :title="$t('dialog.register.cancelRegister')"
                :content="$t('dialog.register.cancelRegisterInform')"/>
    </v-container>
</template>

<script lang="ts">
    import {Component, Vue, Watch, Emit, Ref} from "vue-property-decorator"
    import rules from '@/utils/rules'
    import Confirm from '@/components/dialogs/Confirm.vue';

    // eslint-disable-next-line no-unused-vars
    import {Member} from "@/utils/member";
    // eslint-disable-next-line no-unused-vars
    import IConfirm from '@/components/dialogs/Confirm';

    interface TSVFrom {
        reset(): void

        resetValidation(): void

        validate(): boolean
    }

    @Component({
        name: "Register",
        components: {
            Confirm
        },
        beforeMount() {
            (this as Register).windowChangedHandle()
        }
    })
    export default class Register extends Vue {
        // 닫기 확인 창
        @Ref()
        cancelDialog!: IConfirm
        //
        readonly stepMax = 4
        step = 0
        // 레퍼런스
        @Ref()
        form0!: TSVFrom
        @Ref()
        form1!: TSVFrom
        @Ref()
        form2!: TSVFrom
        @Ref()
        form3!: TSVFrom

        // [메인뷰] 폼 유효성
        formModel = [false, false, false, false]
        jumpActive = [true, false, false, false]

        @Watch('formModel', {deep: true, immediate: true})
        formModelWatcher() {
            Vue.set(this.jumpActive, 1, this.formModel.slice(0, 1).every(value => value))
            Vue.set(this.jumpActive, 2, this.formModel.slice(0, 2).every(value => value))
            Vue.set(this.jumpActive, 3, this.formModel.slice(0, 3).every(value => value))
        }

        @Watch('jumpActive', {deep: true})
        jumpActiveWatcher() {
            if (!this.jumpActive[this.step]) {
                // step은 0과 jumpActive중 가장 큰 인덱스의 위치 사이에 존재해야 하므로
                this.step = Math.max(0, Math.min(this.step, this.jumpActive.filter(value => value).length - 1))
            }
        }

        @Watch('step')
        windowChangedHandle() {
            switch (this.step) {
                case 0:
                    this.prevIcon = 'mdi-cancel'
                    this.nextIcon = 'mdi-chevron-right'
                    // 메세지 세팅
                    this.prevMessage = this.$t('dialog.register.cancelButton').toString()
                    this.nextMessage = this.formModel[0] ? this.nextMessage = this.$t('dialog.register.doNext').toString() : ''
                    // 색 세팅
                    this.prevColor = 'error'
                    this.nextColor = this.formModel[0] ? 'success' : ''
                    break
                case 3:
                    this.prevIcon = 'mdi-chevron-left'
                    this.nextIcon = 'mdi-account-check'
                    this.prevMessage = ''
                    this.nextMessage = this.formModel[3] ? this.nextMessage = this.$t('dialog.register.directLogin').toString() : ''
                    this.prevColor = ''
                    this.nextColor = this.formModel[0] ? 'success' : ''
                    break
                default:
                    this.prevIcon = 'mdi-chevron-left'
                    this.nextIcon = 'mdi-chevron-right'
                    this.prevMessage = ''
                    this.nextMessage = this.formModel[this.step] ? this.nextMessage = this.$t('dialog.register.doNext').toString() : ''
                    this.prevColor = ''
                    this.nextColor = this.formModel[this.step] ? 'success' : ''
                    break
            }
        }

        // [메인뷰] 0 : 동의
        agreeTermOfUse = false
        agreePersonalData = false
        // [메인뷰] 1.이름
        name = ''
        readonly ruleName = [(testee: string) => !!testee || this.$t('dialog.register.emptyName')]
        // [메인뷰] 1.이메일
        email = ''
        emailSearching: string | null = null
        availableEmails: string[] = []
        readonly ruleEmail = [
            (testee: string) => !!testee || this.$t('dialog.register.emptyEmail'),
            (testee: string) => rules.email(testee) || this.$t('dialog.register.invalidEmail'),
        ]

        @Watch('emailSearching')
        emailSearchingWatcher(input: string | null) {
            if (input != null && input.length > 0) {
                let [name, domain] = input.split('@')
                if (typeof domain !== "undefined") {
                    this.email = input
                    this.availableEmails = [input]
                } else {
                    this.availableEmails = (Object.values(this.$t('famous.emailDomains')) as string[]).map(value => {
                        return name + '@' + value
                    })
                }
            } else {
                this.availableEmails = []
            }
        }

        // [메인뷰] 1.전화번호
        phone = ''
        // [메인뷰] 1.주소
        address = ''
        // [메인뷰] 2 : 각종 이벤트 수신여부, 알림 허용여부
        acceptNotice? = false
        acceptAD? = false
        acceptDLog ? = false

        // [하단 네비게이션]
        prevIcon = 'mdi-chevron-left'
        nextIcon = 'mdi-chevron-right'
        prevMessage = ''
        nextMessage = ''
        prevColor = ''
        nextColor = ''


        nextStep() {
            if ((this.step + 1) === this.stepMax) {
                this.onComplete()
                return
            }
            let tmp = Math.min(this.step + 1, this.stepMax - 1)
            if (this.jumpActive[tmp]) {
                this.step = tmp
            }
        }

        prevStep() {
            if (this.step === 0) {
                // this.onCancel()
                this.cancelDialog.open(arg0 => {
                    if (arg0) {
                        this.onCancel()
                    }
                })
                return
            }
            let tmp = Math.max(this.step - 1, 0)
            if (this.jumpActive[tmp]) {
                this.step = tmp
            }
        }

        reset() {
            this.form0?.reset()
            this.form1?.reset()
            this.form2?.reset()
            this.form3?.reset()
        }

        @Emit('cancel')
        onCancel() {
        }

        @Emit('complete')
        onComplete(): Member {
            const dat = {
                name: this.name,
                email: this.email,
                address: this.address,
                phone: this.phone,
                agree: {
                    notice: Boolean(this.acceptNotice),
                    advertisement: Boolean(this.acceptAD),
                    logging: Boolean(this.acceptDLog),
                }
            }
            this.reset()
            return dat
        }


    }
</script>

<style scoped lang="scss">
    .what-is-video {
        width: 100%;
        max-height: 160px;
        object-fit: fill;
    }

    .what-is-text {
        white-space: pre-wrap;
        word-break: keep-all;
    }


</style>

