<template>
  <div>
    <NavbarMobileVue />
    <HomeVideoVue />
    <HomeCarouselVue />
    <LastestPostVue />
    <FooterVue />
  </div>
</template>

<script>
import FooterVue from "../components/front/Footer.vue";
import HomeVideoVue from "../components/front/HomeVideo.vue";
import HomeCarouselVue from "../components/front/HomeCarousel.vue";
import LastestPostVue from "../components/front/LastestPost.vue";
import NavbarMobileVue from "../components/front/NavbarMobile.vue";
import PopularPostVue from "../components/front/PopularPost.vue";

export default {
  name: "IndexPage",
  auth: false,
  head(){
  return{
   title:'Home',
  }
  },
  components: {
    NavbarMobileVue,
    HomeVideoVue,
    HomeCarouselVue,
    PopularPostVue,
    LastestPostVue,
    FooterVue,
  },
};
</script>
