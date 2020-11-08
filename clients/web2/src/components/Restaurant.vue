<template>
    <v-card
            class="mx-auto my-12"
    >
        <v-img
                min-height="100"
                max-height="200"
                src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
                @click="alertName"
        />
        <v-card-title>
            <span @click="alertName">{{name}}</span>
            <v-spacer/>
            <v-rating
                    :value="4.5"
                    color="amber"
                    dense
                    half-increments
                    readonly
                    size="14"/>
        </v-card-title>
        <v-card-text>
            <v-row
                    align="center"
                    class="mx-0"
            >
            </v-row>
            <div class="my-4 subtitle-1">해시태그 들어갈 곳</div>
            <div>Small plates, salads & sandwiches - an intimate setting with 12 indoor seats plus patio seating.</div>
        </v-card-text>
        <v-divider class="mx-4"></v-divider>
        <v-card-actions>
            <v-spacer/>
            <v-btn
                    color="deep-purple lighten-2"
                    text
                    @click="alertName"
            >
                리뷰(100)
            </v-btn>
            <v-btn
                    color="deep-purple lighten-2"
                    text
                    @click="alertName"
            >
                댓글(100)
            </v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
    import Vue from "vue"

    export default Vue.extend({
        name: "restaurant-card",
        props: {
            name: String,
            avatar: [URL, String],
        },
        methods: {
            alertName(){
                alert(this.$props.name)
            },
        }
    })
</script>

<style scoped>
    .Restaurant {
        min-width: 280px;
        max-width: 400px;
        height: 300px;
    }
</style>