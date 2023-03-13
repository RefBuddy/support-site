<template>
    <section class="screenshots">
        <div class="screenshots__container">
            <button
                class="screenshots__arrow screenshots__arrow--left"
                @click="prevImage"
            >
                &#60;
            </button>
            <div class="screenshots__items" ref="screenshotsItems">
                <img
                    v-if="currentImage === 0"
                    class="screenshot__item"
                    src="@/assets/Stats.png"
                    alt="Stats"
                />
                <img
                    v-if="currentImage === 1"
                    class="screenshot__item"
                    src="@/assets/Assignor.png"
                    alt="Assignor"
                />
                <img
                    v-if="currentImage === 2"
                    class="screenshot__item"
                    src="@/assets/Game.png"
                    alt="Game"
                />
            </div>
            <button
                class="screenshots__arrow screenshots__arrow--right"
                @click="nextImage"
            >
                &#62;
            </button>
        </div>
    </section>
</template>

<script>
export default {
    data() {
        return {
            currentImage: 0,
        };
    },
    methods: {
        nextImage() {
            this.currentImage = (this.currentImage + 1) % 3;
        },
        prevImage() {
            this.currentImage = (this.currentImage + 2) % 3;
        },
    },
};
</script>
