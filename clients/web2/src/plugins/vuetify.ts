import Vue from 'vue';
import Vuetify, {colors} from 'vuetify/lib';

Vue.use(Vuetify);
console.log(colors.orange.lighten3)
export default new Vuetify({
    theme: {
        themes : {
            // light : {
            //     primary: '#ffcc80',
            //     success: '#3cd1c2',
            //     info: '#ffaa2c',
            //     error: '#f83e70'
            // },
        }
    }
});
