<template>
<div>
    <v-carousel v-model="model" height="auto">
        <v-carousel-item>
            <div class="bg_image" style="background-image: url('http://localhost:8800/images/dhaka_2.jpg');">
                <div class="w__40">
                    <h1 class="text--white">Welcome to my Nuxt and laravel website</h1>
                    <p class="text--white">
                        Lorem ipsum dolor, sit amet consectetur adipisicing elit. Illo ut
                        excepturi, incidunt facilis temporibus perferendis, veritatis,
                        sapiente molestias quibusdam aspernatur consequatur fugiat quasi
                        dicta nesciunt provident amet? Officia, exercitationem laboriosam.
                    </p>
                </div>
            </div>
        </v-carousel-item>
        <v-carousel-item>
            <div class="bg_image" style="
            background-image: url('http://localhost:8800/images/dhaka_1.jpg');
          ">
                <div class="w__40">
                    <h1 class="text--white">Welcome to my Nuxt a  nd laravel website</h1>
                    <p class="text--white">
                        Lorem ipsum dolor, sit amet consectetur adipisicing elit. Illo ut
                        excepturi, incidunt facilis temporibus perferendis, veritatis,
                        sapiente molestias quibusdam aspernatur consequatur fugiat quasi
                        dicta nesciunt provident amet? Officia, exercitationem laboriosam.
                    </p>
                </div>
            </div>
        </v-carousel-item>
    </v-carousel>
</div>
</template>

<script>
export default {
    name: "HeroCamp",
    data: () => ({
        model: 0,
    }),
};
</script>

<style>
</style>
