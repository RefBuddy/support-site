<template>
    <div class="ref-buddy-page">
        <header class="header">
            <img
                src="@/assets/Ref_Buddy.jpg"
                class="header__image"
                alt="Ref Buddy logo"
            />
            <h1 class="header__title">Ref Buddy</h1>
        </header>
        <main class="main">
            <section class="introduction">
                <p class="introduction__description">
                    Ref Buddy is a revolutionary mobile scheduling solution
                    designed specifically for hockey officials and assignors.
                    With its intuitive interface and powerful features, Ref
                    Buddy makes it easy for assignors to schedule the right
                    officials for each game and for officials to manage their
                    schedules and stay informed about their assignments.
                </p>
                <p class="introduction__description">
                    Each official has a personal profile where they can access
                    season stats, as well as view and update their availability.
                    Officials can also see their upcoming games and review
                    previous games. All games have their box scores linked,
                    making it easier for officials to prepare and stay
                    organized.
                </p>
                <p class="introduction__description">
                    With Ref Buddy, assignors have access to all the information
                    they need in one convenient place, including official
                    profiles, league schedule, and real-time availability
                    updates.
                </p>
                <p class="introduction__description">
                    Streamline your hockey official scheduling with Ref Buddy,
                    the ultimate mobile scheduling solution for elite hockey
                    leagues.
                </p>
            </section>
            <section class="features">
                <h2 class="features__title">Features</h2>
                <ul class="features__list">
                    <li>Automatically imported league schedule</li>
                    <li>Assign referees and linesmen to games</li>
                    <li>Individual official profiles</li>
                    <li>Stats for every official</li>
                    <li>Quick access to assigned games and box scores</li>
                    <li>
                        When an official gets assigned a game, they receive a
                        notification from Ref Buddy on their phone
                    </li>
                    <li>Send alerts and messages to officiating staff</li>
                    <li>Game reminder system</li>
                    <li>Incident reporting system</li>
                    <li>Hockey Canada rulebook</li>
                    <li>League rulebook</li>
                    <li>More to come!</li>
                </ul>
            </section>
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
            <footer ref="footer" class="footer">
                <p>
                    Is your league interested in using Ref Buddy to assign games
                    and track stats? Feel free to reach out with any questions
                    and we will get back to you right away!
                </p>
                <form ref="footerForm" class="footer__form">
                    <input v-model="form.Name" type="text" placeholder="Name" />
                    <input
                        v-model="form.Email"
                        type="email"
                        placeholder="Email"
                    />
                    <textarea
                        v-model="form.Message"
                        placeholder="Message"
                    ></textarea>
                    <button @click.prevent="submitForm">Submit</button>
                </form>
            </footer>
            <div v-if="success" class="footer__success">
                We have received your message and will get back to you as soon
                as possible!
            </div>
        </main>
    </div>
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            currentImage: 0,
            form: {
                Name: "",
                Email: "",
                Message: "",
            },
            success: false,
        };
    },
    methods: {
        async submitForm() {
            try {
                await axios.post(
                    "https://us-central1-ref-buddy-d7be3.cloudfunctions.net/sendContactForm",
                    { data: this.form }
                );
                this.success = true;
                this.$refs.footerForm.style.display = "none";
                this.$refs.footer.style.display = "none";
            } catch (error) {
                console.error(error);
                alert("An error occurred while submitting the form.");
            }
        },
        nextImage() {
            this.currentImage = (this.currentImage + 1) % 3;
        },
        prevImage() {
            this.currentImage = (this.currentImage + 2) % 3;
        },
    },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=VT323&display=swap");

body {
    font-family: "VT323", monospace;
    background-color: rgba(255, 145, 0, 0.73);
    font-size: 20px;
}

.ref-buddy-page {
    width: auto;
    max-width: 550px;
    min-width: 300px;
    height: 92%;
    min-height: 450px;
    margin: 0 auto;
    padding-top: 20px;
    padding-left: 50px;
    padding-right: 50px;
    background-color: white;
    border-radius: 0;
    border: 3px solid black;
    position: absolute;
    top: calc(50% - 10px);
    left: 50%;
    transform: translate(-50%, -50%);
    box-shadow: 15px 15px 0px 0px black;
}

.header {
    height: 70px;
    width: auto;
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    position: sticky;
    top: 0;
    left: 0;
    right: 0;
    z-index: 99;
    margin: 0 -16px 16px -16px;
}

.header__image {
    width: 70px;
    height: 70px;
    margin-right: 16px;
    float: left;
    border-radius: 12px;
}

.header__title {
    margin: 0;
    padding-top: 15px;
    margin-right: 15px;
}

.main {
    flex: 1;
    padding: 16px;
    width: auto;
    max-width: 550px;
    height: 78%;
    overflow-y: scroll;
    margin-left: 10px;
}

@media screen and (max-height: 760px) {
    .main {
        height: 75%;
    }
}

@media screen and (max-height: 660px) {
    .ref-buddy-page {
        height: 85%;
    }
    .main {
        height: 70%;
    }
}

@media screen and (max-height: 595px) {
    .main {
        height: 65%;
    }
}

@media screen and (max-width: 480px) {
    .ref-buddy-page {
        background-color: transparent;
        border: none;
        box-shadow: none;
    }
    .main {
        height: 80%;
        margin-top: -15px;
    }
    .header {
        margin: -20px -8px 0 -8px;
    }
}

.main::-webkit-scrollbar {
    display: none;
}

.introduction {
    width: calc(100% - 7px);
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    box-shadow: 5px 5px 0px 0px black;
    margin: -4px -16px 26px -16px;
}

.features {
    width: calc(100% - 7px);
    background-color: white;
    border: 1px solid black;
    padding: 16px;
    border-radius: 0;
    box-shadow: 5px 5px 0px 0px black;
    margin: -4px -16px 26px -16px;
}

.features__title {
    color: black;
    width: auto;
    max-width: 125px;
    padding: 10px;
    margin: 0 auto;
    text-align: center;
}

.features-list {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 16px;
}

.features-list__item {
    display: flex;
    gap: 8px;
}

.features-list__icon {
    width: 24px;
    height: 24px;
    object-fit: contain;
}

.features-list__text {
    font-weight: bold;
}

.screenshots {
    width: calc(100% - 7px);
    padding: 16px;
    border-radius: 0;
    text-align: center;
    margin: 0 -16px 26px -16px;
    overflow-x: auto;
}

.screenshots__title {
    color: black;
    border: 1px solid orange;
    width: auto;
    max-width: 225px;
    padding: 10px;
    margin: 0 auto;
    text-align: center;
}

.screenshots__items {
    display: flex;
    flex-wrap: nowrap;
    overflow-x: auto;
    border: 1px solid rgba(175, 175, 175, 0.236);
    border-radius: 12px;
    box-shadow: 5px 5px 5px 0px rgba(175, 175, 175, 0.236);
}

.screenshots__items > * {
    flex-shrink: 0;
    width: 100%;
    max-width: 100%;
    margin-right: 16px;
}

.screenshots__items > :last-child {
    margin-right: 0;
}

.screenshots__container {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    margin: 0 -16px;
}

.screenshots__arrow {
    background-color: transparent;
    border: none;
    font-size: 32px;
    padding: 8px 16px;
    cursor: pointer;
}

.screenshots__arrow--left {
    margin-right: 16px;
}

.screenshots__arrow--right {
    margin-left: 16px;
}

.footer {
    width: calc(100% - 7px);
    min-height: 245px;
    background-color: white;
    padding: 16px;
    border: 1px solid black;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 -16px 16px -16px;
}

.footer__success {
    max-height: 70px;
    width: 80%;
    display: flex;
    padding: 16px;
    font-size: 20px;
    background-color: white;
    border: 1px solid black;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 auto;
}

.footer__form {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 16px;
}

.footer__form > * {
    width: 100%;
    margin-bottom: 16px;
}

.footer__form > textarea {
    height: 100px;
}

.footer__form > button {
    background-color: black;
    color: white;
    border: none;
    padding: 8px 16px;
    cursor: pointer;
}
</style>
