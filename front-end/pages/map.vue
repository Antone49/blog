<template>
  <div>
    <NavbarMobileVue />
    <AllPostHeroVue />
    <MapVue />
    <FooterVue />
  </div>
</template>

<script>
import MapVue from "../components/front/Map.vue";
import AllPostHeroVue from "../components/front/AllPostHero.vue";
import FooterVue from "../components/front/Footer.vue";
import NavbarMobileVue from "../components/front/NavbarMobile.vue";

export default {
  name: "MapPage",
  auth: false,
  head(){
  return{
   title:'Blog',
  }
  },
  components: {
    NavbarMobileVue,
    FooterVue,
    MapVue,
    AllPostHeroVue,
  },
};
</script>
