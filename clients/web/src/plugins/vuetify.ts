import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
            dark: {},
            light: {
                primary: process.env.VUE_APP_THEME_PRIMARY,
                secondary: '#4287F5'
            },
        }
    }
});
