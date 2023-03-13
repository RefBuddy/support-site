<template>
    <footer ref="footer" class="footer">
        <p>
            Is your league interested in using Ref Buddy to assign games and
            track stats? Feel free to reach out with any questions and we will
            get back to you right away!
        </p>
        <form ref="footerForm" class="footer__form">
            <input v-model="form.Name" type="text" placeholder="Name" />
            <input v-model="form.Email" type="email" placeholder="Email" />
            <input v-model="form.League" type="text" placeholder="League" />
            <textarea v-model="form.Message" placeholder="Message"></textarea>
            <button @click.prevent="submitForm">Submit</button>
        </form>
    </footer>
    <div v-if="success" class="footer__success">
        We have received your message and will get back to you as soon as
        possible!
    </div>
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            form: {
                Name: "",
                Email: "",
                League: "",
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
    },
};
</script>
