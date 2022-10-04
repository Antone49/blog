<template>
<v-app-bar app :color="bg" inverted-scroll scroll-threshold="150">
    <NavbarContentVue />
</v-app-bar>
</template>

<script>
import NavbarContentVue from "./NavbarContent.vue";

export default {
    name: "NavbarMobile",
    components: {
        NavbarContentVue,
    },
    data: () => ({
        showHeader: false,
        nav: true,
        bg: 'grey'
    }),
    mounted() {
        window.onscroll = () => {
            this.changeColor();
        };
    },
    methods: {
        changeColor() {
            if (document.body.scrollTop > 250 ||
                document.documentElement.scrollTop > 250) {
                this.bg = 'grey';
            } else {
                this.bg = 'transparent';
            }
        },
    },
};
</script>

<style>
.is-hidden {
    transform: translateY(-100%);
}
</style>
