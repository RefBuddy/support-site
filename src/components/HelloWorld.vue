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
                    Ref Buddy is a mobile scheduling system used to assign
                    on-ice officials to hockey games.
                </p>
                <p class="introduction__description">
                    Each official gets their own profile where they can see
                    their season stats, view and update availability, see
                    upcoming games, and view previous games.
                </p>
                <p class="introduction__description">
                    Ref Buddy allows for assignors to quickly and seemlessly add
                    officials to games, as well as view everything in their
                    profile.
                </p>
            </section>
            <section class="features">
                <h2 class="features__title">Features</h2>
                <ul class="features__list">
                    <li>Assign referees and linesmen to games</li>
                    <li>Automatically imported league schedule</li>
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
                        @click="scrollLeft"
                    >
                        &#60;
                    </button>
                    <div class="screenshots__items" ref="screenshotsItems">
                        <img
                            class="screenshot__item"
                            src="@/assets/Stats.png"
                            alt="Screenshot 1"
                        />
                        <img
                            class="screenshot__item"
                            src="@/assets/logo.png"
                            alt="Screenshot 2"
                        />
                    </div>
                    <button
                        class="screenshots__arrow screenshots__arrow--right"
                        @click="scrollRight"
                    >
                        &#62;
                    </button>
                </div>
            </section>
            <footer class="footer">
                <h2 class="footer__title">Contact Us</h2>
                <form class="footer__form">
                    <input type="text" placeholder="Name" />
                    <input type="email" placeholder="Email" />
                    <textarea placeholder="Message"></textarea>
                    <button type="submit">Submit</button>
                </form>
            </footer>
        </main>
    </div>
</template>

<script>
export default {
    data() {
        return {
            currentIndex: 0,
        };
    },
    mounted() {
        this.screenshotItems = this.$refs.screenshotsItems.children;
    },
    methods: {
        scrollLeft() {
            if (this.currentIndex > 0) {
                this.currentIndex--;
                for (let i = 0; i < this.screenshotItems.length; i++) {
                    this.screenshotItems[i].style.transform = `translateX(-${
                        this.currentIndex * this.screenshotItems[i].offsetWidth
                    }px)`;
                }
            }
        },
        scrollRight() {
            if (this.currentIndex < this.screenshotItems.length - 1) {
                this.currentIndex++;
                for (let i = 0; i < this.screenshotItems.length; i++) {
                    this.screenshotItems[i].style.transform = `translateX(-${
                        this.currentIndex * this.screenshotItems[i].offsetWidth
                    }px)`;
                }
            }
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
    height: 80%;
    min-height: 300px;
    margin: 0 auto;
    padding: 40px;
    background-color: white;
    border-radius: 0;
    border: 3px solid black;
    position: absolute;
    top: 50%;
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
    height: 80%;
    overflow-y: scroll;
    margin-left: 10px;
}

@media screen and (max-height: 690px) {
    .ref-buddy-page {
        height: 70%;
    }
    .main {
        height: 70%;
    }
}

@media screen and (max-height: 540px) {
    .ref-buddy-page {
        height: 60%;
    }
    .main {
        height: 60%;
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
    background-color: white;
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
    box-shadow: 10px 10px 0px 0px rgba(175, 175, 175, 0.236);
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
    background-color: white;
    padding: 16px;
    border: 1px solid black;
    border-radius: 0;
    text-align: center;
    box-shadow: 5px 5px 0px 0px black;
    margin: 0 -16px 16px -16px;
}

.footer__title {
    color: black;
    margin-bottom: 16px;
}

.footer__form {
    display: flex;
    flex-direction: column;
    align-items: center;
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
