<template>
    <v-container
    >
        <v-row>
            <v-col v-for="k in cache" :key="k.value.query.key()"
                   xl="2"
                   lg="3"
                   md="4"
                   sm="6"
                   xs="12"
            >
                <Restaurant v-bind:name="k.value.name" v-bind:avatar="k.name"/>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import sddap from "@/store/sddapp"
    import Restaurant from "@/components/Restaurant"
    import {ServiceType, TypeQuery} from "@/classes/ddapp/query";
    import {mapState} from "vuex";

    import {Component, Vue} from 'vue-property-decorator';

    @Component({
        components: {
            Restaurant,
        },
        beforeCreate() {
            sddap.dispatch('reload', {
                request: new TypeQuery(ServiceType.Restaurant)
            })
        },
        computed: {
            ...mapState({
                cache: 'cached'
            }),
        },
    })
    export default class FoodGrid extends Vue {

    }
</script>
<style scoped>
</style>