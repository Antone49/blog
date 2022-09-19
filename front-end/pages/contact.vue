<template>
  <div>
    <NavbarMobileVue />
    <ContactHeroVue />
    <ContactFormVue />
    <FooterVue />
  </div>
</template>
<script>
import ContactFormVue from "../components/front/ContactForm.vue";
import ContactHeroVue from "../components/front/ContactHero.vue";
import FooterVue from "../components/front/Footer.vue";
import NavbarMobileVue from "../components/front/NavbarMobile.vue";

export default {
  name: "ContactPage",
  head(){
  return{
   title:'Contact',
  }
  },
  components: {
    ContactHeroVue,
    FooterVue,
    NavbarMobileVue,
    ContactFormVue,
  },
};
</script>
<style >
</style>
