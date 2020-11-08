<template>
    <v-dialog v-model="isDialogOpen" max-width="800">
        <!-- 다이얼로그 열기 버튼 -->
        <template v-slot:activator="{on, attrs}">
            <slot name="activator" v-bind:on="on" v-bind:attrs="attrs"></slot>
        </template>
        <v-card>
            <v-window v-model="step" vertical>
                <!-- 로그인 요청-->
                <v-window-item :value="0">
                    <v-container fluid style="padding-top: 0;padding-bottom: 0" class="primary">
                        <v-row>
                            <v-col class="d-flex" :class="{'mobile-logo':this.$vuetify.breakpoint.smAndDown}" cols="12"
                                   md="7" align-self="center">
                                <v-img class="align-self-center" :src="require('@/assets/MakItGo_transparent.png')"/>
                            </v-col>
                            <v-col cols="12" md="5" style="background-color: white">
                                <v-list>
                                    <v-list-item-title class="d-flex justify-center"><h2>
                                        {{$t('dialog.login.loginMessage')}}</h2></v-list-item-title>
                                    <v-divider/>
                                    <!-- 로그인 -->
                                    <v-list-item-content class="pt-6">
                                        <v-form>
                                            <v-text-field
                                                    prepend-inner-icon="mdi-account"
                                                    filled rounded clearable dense
                                                    v-model="formID"
                                            />
                                            <v-text-field
                                                    prepend-inner-icon="mdi-lock"
                                                    filled rounded clearable dense
                                                    :type="showPassword ? 'text' : 'password'"
                                                    :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                                                    @click:append="showPassword = !showPassword"
                                                    v-model="formPW"

                                            />
                                            <v-btn block rounded dark depressed color="success">
                                                {{$t('dialog.login.loginButton')}}
                                            </v-btn>
                                        </v-form>
                                    </v-list-item-content>
                                    <v-list-item-subtitle class="d-flex justify-center">
                                        <v-btn
                                                block small text color="primary"
                                                @click="step=1">
                                            {{$t('dialog.login.notMember')}}
                                        </v-btn>
                                    </v-list-item-subtitle>
                                    <v-list-item-subtitle class="d-flex justify-center">
                                        <v-btn
                                                block small text color="primary"
                                                @click="step=2">
                                            {{$t('dialog.login.accountForgot')}}
                                        </v-btn>
                                    </v-list-item-subtitle>
                                    <!-- OAuth2 로그인 -->
                                    <v-divider class="mb-2"/>
                                    <v-list-item-subtitle class="d-flex justify-center">
                                        {{$t('dialog.login.loginWithOAuth')}}
                                    </v-list-item-subtitle>
                                    <v-list-item class="d-flex justify-center">
                                        <v-layout>
                                            <v-flex>
                                                <v-btn icon :color="es_google_primary" x-large>
                                                    <v-icon>mdi-google</v-icon>
                                                </v-btn>
                                            </v-flex>
                                            <v-flex>

                                                <v-btn icon :color="es_facebook_primary" x-large>
                                                    <v-icon>mdi-facebook</v-icon>
                                                </v-btn>
                                            </v-flex>
                                            <v-flex>
                                                <v-btn icon :color="es_twitter_primary" x-large>
                                                    <v-icon>mdi-twitter</v-icon>
                                                </v-btn>
                                            </v-flex>
                                            <v-flex>
                                                <v-btn icon :color="es_linkedin_primary" x-large>
                                                    <v-icon>mdi-linkedin</v-icon>
                                                </v-btn>
                                            </v-flex>
                                            <v-flex>
                                                <v-btn icon :color="es_microsoft_primary" x-large>
                                                    <v-icon>mdi-microsoft</v-icon>
                                                </v-btn>
                                            </v-flex>
                                        </v-layout>
                                    </v-list-item>
                                    <!-- 호스팅 로그인 -->
                                </v-list>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-window-item>
                <!-- 회원가입 -->
                <v-window-item :value="1">
                    <Register @cancel="step=0" @complete="registered"/>
                </v-window-item>
                <!-- 비밀번호 찾기 -->
                <v-window-item :value="2">
                    <div class="d-flex flex-column align-start justify-center pa-4" style="height: 60vh">
                        <p class="my-auto mx-auto">{{$t('placeholder[0]')}}</p>
                        <v-btn class="mb-4 mx-auto" width="100%" @click="step=0" color="error">죄송합니다...</v-btn>
                    </div>
                </v-window-item>
            </v-window>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
    import {Vue, Component} from "vue-property-decorator"
    import {mapState} from "vuex"
    import Register from "./Register.vue"
    // eslint-disable-next-line no-unused-vars
    import {Member} from "@/utils/member";

    @Component({
        name: "Login.vue",
        components: {
            Register
        },
        computed: {
            ...mapState("login", ["supportAuth"]
            ),
        },

    })
    export default class LoginDialog extends Vue {
        isDialogOpen = false
        step = 0
        formID = ""
        formPW = ""
        showPassword = false

        //
        registered(m: Member) {
            console.log(m)
            this.step = 0

        }

        //
        es_google_primary = process.env.VUE_APP_EXTERN_SERVICE_GOOGLE_PRIMARY_COLOR
        es_facebook_primary = process.env.VUE_APP_EXTERN_SERVICE_FACEBOOK_PRIMARY_COLOR
        es_twitter_primary = process.env.VUE_APP_EXTERN_SERVICE_TWITTER_PRIMARY_COLOR
        es_linkedin_primary = process.env.VUE_APP_EXTERN_SERVICE_LINKEDIN_PRIMARY_COLOR
        es_microsoft_primary = process.env.VUE_APP_EXTERN_SERVICE_MICROSOFT_PRIMARY_COLOR
    }
</script>

<style scoped lang="scss">
    .mobile-logo {
        @media(max-height: 800px) {
            max-height: 9.5rem;
        }
        @media(max-height: 800px) and (min-width: 400px) {
            max-height: 13.5rem;
        }
        @media(min-height: 800px) {
            max-height: 18.5rem;
        }
    }
</style>