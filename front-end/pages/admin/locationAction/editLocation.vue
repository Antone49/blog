<template>
<div>
    <SidebarVue />
    <v-container>
        <v-col xl="12" lg="12" md="12" sm="12" xs="12">
            <nuxt-link to="/admin/location">
                <v-btn color="secondary" class="mb-4">Annuler</v-btn>
            </nuxt-link>
        </v-col>
        <v-card>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <v-card-title>Edit location</v-card-title>
                <v-card-subtitle>
                    <form action="" class="form">
                        <v-row>
                            <v-col xl="6" lg="6" md="6" sm="12" xs="12">
                                <v-text-field label="Nom" class="form__control" v-model="name" hide-details="auto" />
                            </v-col>
                            <br>
                            <v-col xl="6" lg="6" md="6" sm="12" xs="12">
                                <v-select :items="items" label="Image" v-on:change="onMarkerChangeIcon" v-model="markerImage">
                                    <template v-slot:selection="{ item, index }">
                                        <img :src="'/images/location/' + item" width="20" height="20">
                                    </template>
                                    <template v-slot:item="{ item }">
                                        <img :src="'/images/location/' + item" width="30" height="30">
                                    </template>
                                </v-select>
                            </v-col>
                        </v-row>

                        <v-btn @click="saveBtn" color="error" class="login__btn">Sauvegarder</v-btn>
                    </form>
                </v-card-subtitle>
            </v-col>
            <v-col xl="12" lg="12" md="12" sm="12" xs="12">
                <h1 class="m-3">Map</h1>
                <div id="map" style="height: 90vh"></div>
            </v-col>
        </v-card>
    </v-container>
</div>
</template>

<script>
import SidebarVue from "/components/admin/Sidebar.vue";

import {
    getLocation,
    addLocation,
    updateLocation
} from '/functions/location.js'

export default {
    name: "EditLocationPage",
    head() {
        return {
            title: 'EditLocation',
        }
    },
    components: {
        SidebarVue,
    },
    data() {
        return {
            id: null,
            name: '',
            mapCenter: [23.791474915526848, 90.41672529214094],
            mapZoom: 15,
            marker: null,
            markerImage: './default.png',
            items: [], 

        }
    },
    methods: {
        saveBtn() {
            if (this.name.trim().length === 0) {
                this.$toast.warning("Veuillez saisir un nom !");
            } else {
                if (this.id == 0) {
                    addLocation(this.$auth.strategy.token.get(), this.name, this.markerImage, this.marker.getLatLng().lng, this.marker.getLatLng().lat)
                } else {
                    updateLocation(this.$auth.strategy.token.get(), this.id, this.name, this.markerImage, this.marker.getLatLng().lng, this.marker.getLatLng().lat)
                }

                this.$router.push({
                    path: '/admin/location'
                })
            }
        },
        onMarkerChangeIcon() {
            this.markerChangeIcon(this.markerImage)
        },
        markerChangeIcon(icon) {
            var myIcon = L.icon({
                iconUrl: "/images/location/" + icon,
                iconSize: [25, 25]
            });
            this.marker.setIcon(myIcon)
        },
        importAllImages(r) {
            r.keys().forEach(key => (
                this.items.push(key)
            ));
        },
        loadMap() {
            // Creating map options
            var mapOptions = {
                center: this.mapCenter,
                zoom: this.mapZoom
            }
            // Creating a map object
            var map = new L.map('map', mapOptions);

            // Creating a Layer object
            var layer = new L.TileLayer('https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}.png');

            // Adding layer to the map
            map.addLayer(layer);

            // Creating Marker Options
            var markerOptions = {
                draggable: true
            }
            // Creating a Marker
            this.marker = L.marker(this.mapCenter, markerOptions);

            this.markerChangeIcon(this.markerImage)

            // Adding marker to the map
            this.marker.addTo(map);
        },
    },
    mounted: function () {
        this.id = parseInt(this.$route.query.id)

        this.loadMap()

        // load item
        this.importAllImages(require.context('/static/images/location/', true, /\.png$/));

        if (this.id != 0) {
            getLocation(this.$auth.strategy.token.get(), this.id).then(result => {
                if (result != null) {
                    this.name = result.name
                    this.marker.setLatLng([result.latitude, result.longitude])

                    this.markerImage = result.image
                    this.markerChangeIcon(this.markerImage)
                }
            });
        }
    },
};
</script>

<style>
</style>
