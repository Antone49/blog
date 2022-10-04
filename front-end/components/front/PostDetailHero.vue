<template>
<div>
    <div class="bg_image" :style="'background-image: url(' + image + ');height:70vh;q'">
        <NavbarContentVue />
    </div>
</div>
</template>

<script>
import NavbarContentVue from "./NavbarContent.vue";

export default {
    name: "PostDetailHero",
    props: {
        image: String,
    },
    components: {
        NavbarContentVue,
    },
    data: () => ({
        model: 0,
    }),
};
</script>

<style>
</style>
