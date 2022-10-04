<template>
<div>
    <div class="bg_image" style="background-image: url('http://localhost:8800/images/dhaka_2.jpg');height:60vh;">
        <NavbarContentVue />
    </div>
</div>
</template>

<script>
import NavbarContentVue from "./NavbarContent.vue";

export default {
    name: "ContactHero",
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
