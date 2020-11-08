<template>
    <v-app-bar
            app
            dark
            prominent
            src="img/soul-food-thai-hong-kong.jpg"
            fade-img-on-scroll
            shrink-on-scroll
            color="primary"
    >
        <template v-slot:img="{ props }">
            <v-img v-bind="props" gradient="to top right, rgba(100,115,201,.7), rgba(25,32,72,.7)"></v-img>
        </template>
        <div
                class="d-flex align-center"
                style="cursor:pointer"
                @click="linkHome()"
        >
            <v-avatar
                    alt="Vuetify Logo"
                    class="shrink mr-2"
                    contain
                    transition="scale-transition"
                    width="50"
            ><img src="@/assets/logo.png"></v-avatar>
            <span class="text--white" style="font-family: MaplestoryOTFBold;letter-spacing: 1px">Quick Food</span>
        </div>
        <v-spacer/>
        <transition name="fade"><v-btn v-if="isShowToTop" text @click="scrollTop"><v-icon size="36">mdi-chevron-up-circle-outline</v-icon></v-btn></transition>
        <v-spacer/>
        <v-btn
                href="https://github.com/vuetifyjs/vuetify/releases/latest"
                target="_blank"
                text
        >
            <span class="mr-2">Latest Release</span>
            <v-icon>mdi-open-in-new</v-icon>
        </v-btn>
        <template v-slot:extension>
            <v-tabs fixed-tabs>
                <v-tab :to="environ.VUE_APP_SHOPS_URL"><h2>{{environ.VUE_APP_SHOPS_TABNAME}}</h2></v-tab>
                <v-tab :to="environ.VUE_APP_FAVOR_URL"><h2>{{environ.VUE_APP_FAVOR_TABNAME}}</h2></v-tab>
                <v-tab :to="environ.VUE_APP_ORDER_URL"><h2>{{environ.VUE_APP_ORDER_TABNAME}}</h2></v-tab>
            </v-tabs>
        </template>
    </v-app-bar>
</template>

<script>
    import Vue from 'vue'
    import router from '@/router/index'

    export default Vue.extend({
        name: "TopBar",
        data(){
            return{
                isShowToTop : false,
            }
        },
        computed: {
            environ() {
                return process.env
            }
        },
        methods: {
            linkHome: () => {
                router.push("/");
            },
            scrollTop() {
                // this.$smoothScroll({})
                window.scrollTo({
                    top : 0,
                    behavior : 'smooth',
                })
            },
            scrollHandle() {
                this.isShowToTop = window.scrollY!== 0
            },
        },
        mounted(){
            window.addEventListener('scroll', this.scrollHandle)
        },
        beforeDestroyed(){
            window.removeEventListener('scroll', this.scrollHandle)
        }
    })
</script>

<style scoped>
    @font-face {
        font-family: 'MaplestoryOTFBold';
        src: url('https://cdn.jsdelivr.net/gh/projectnoonnu/noonfonts_20-04@2.1/MaplestoryOTFBold.woff') format('woff');
        font-weight: normal;
        font-style: normal;
    }
    .fade-enter {
        opacity: 0;
    }
    .fade-enter-active {
        transition: opacity 1s;
        opacity: 1;
    }

    .fade-leave {

    }
    .fade-leave-active {
        transition: opacity 1s;
        opacity: 0;
    }
</style>