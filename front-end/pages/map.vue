<template>
  <div>
    <NavbarVue />
    <AllPostHeroVue />
    <MapVue />
    <FooterVue />
  </div>
</template>

<script>
import MapVue from "../components/front/Map.vue";
import AllPostHeroVue from "../components/front/AllPostHero.vue";
import FooterVue from "../components/front/Footer.vue";
import NavbarVue from "../components/front/Navbar.vue";

export default {
  name: "BlogPage",
  head(){
  return{
   title:'Blog',
  }
  },
  components: {
    NavbarVue,
    FooterVue,
    MapVue,
    AllPostHeroVue,
  },
};
</script>
