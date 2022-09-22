<template>
<div>
  <div class="bg_image" style="background-image: url('/images/dhaka_2.jpg');height:60vh;">
        <NavbarContentVue />
    </div>
</div>
</template>

<script>
import NavbarContentVue from "./NavbarContent.vue";

export default {
    name: "PostDetailHero",
    props: {
        title: String,
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
