<template>
    <v-dialog v-model="model" persistent max-width="400">
        <v-card>
            <v-card-title class="headline grey lighten-2">{{title}}</v-card-title>
            <v-card-text>{{content}}</v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                        color="error"
                        text
                        @click="cancelHandle"
                >{{cancel ? cancel : $t('word.cancel')}}
                </v-btn>
                <v-btn
                        color="success"
                        text
                        @click="confirmHandle"
                >{{confirm ? confirm : $t('word.confirm')}}
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
    import Vue from 'vue'
    import Component from 'vue-class-component'
    import {Prop, Emit} from 'vue-property-decorator'

    @Component({
        name: 'Confirm',
    })
    export default class Confirm extends Vue {
        @Prop({
            type: String,
            required: true
        })
        title!: string

        @Prop({
            type: String,
            default: ''
        })
        content!: string


        @Prop({
            type: String
        })
        cancel?: string

        @Prop({
            type: String,
            default: ''
        })
        confirm?: string

        @Emit('cancel')
        onCancel() {
        }

        @Emit('confirm')
        onConfirm() {
        }

        //
        model = false
        private temporal?: (arg0: boolean) => void

        //
        open(handle?: (arg0: boolean) => void) {
            this.model = true
            this.temporal = handle
        }

        private cancelHandle() {
            this.temporal?.(false)
            this.temporal = undefined
            this.onCancel()
            this.model = false
        }

        private confirmHandle() {
            this.temporal?.(true)
            this.temporal = undefined
            this.onConfirm()
            this.model = false
        }
    }
</script>

<style scoped>

</style>