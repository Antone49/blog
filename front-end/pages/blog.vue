<template>
  <div>
    <NavbarMobileVue />
    <AllPostHeroVue />
    <AllPostVue />
    <FooterVue />
  </div>
</template>

<script>
import AllPostVue from "../components/front/AllPost.vue";
import AllPostHeroVue from "../components/front/AllPostHero.vue";
import FooterVue from "../components/front/Footer.vue";
import NavbarMobileVue from "../components/front/NavbarMobile.vue";

export default {
  name: "BlogPage",
  head(){
  return{
   title:'Blog',
  }
  },
  components: {
    NavbarMobileVue,
    FooterVue,
    AllPostVue,
    AllPostHeroVue,
  },
};
</script>
