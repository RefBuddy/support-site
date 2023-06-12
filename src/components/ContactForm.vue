<template>
  <footer ref="footer" class="footer">
    <h2 class="component_title">Join Ref Buddy</h2>
      <p v-html=joinText></p>
    <form ref="footerForm" class="footer__form">
      <div class="form-group">
        <div class="input-wrapper">
          <input v-model="form.Name" type="text" placeholder="Name" />
          <div v-if="errors.Name" class="error">{{ errors.Name }}</div>
        </div>
      </div>
      <div class="form-group">
        <div class="input-wrapper">
          <input v-model="form.Email" type="email" placeholder="Email" />
          <div v-if="errors.Email" class="error">{{ errors.Email }}</div>
        </div>
      </div>
      <div class="form-group">
        <div class="input-wrapper">
          <input v-model="form.League" type="text" placeholder="League" />
          <div v-if="errors.League" class="error">{{ errors.League }}</div>
        </div>
      </div>
      <div class="form-group">
        <div class="input-wrapper">
          <textarea v-model="form.Message" placeholder="Message"></textarea>
          <div v-if="errors.Message" class="error">{{ errors.Message }}</div>
        </div>
      </div>
      <button @click.prevent="submitForm">Submit</button>
    </form>
  </footer>
  <div v-if="success" class="footer__success">
    We have received your message and will get back to you as soon as possible!
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
      errors: {},
      success: false,
      joinText: "Fill out the following form to find out how your league can start using Ref Buddy—the ultimate mobile solution for elite hockey leagues."
    };
  },
  methods: {
    async submitForm() {
        // Clear any existing error messages
        this.errors = {};

        // Check that all fields have a value
        if (!this.form.Name) {
            this.errors.Name = "Please enter your name.";
        }
        if (!this.form.Email) {
            this.errors.Email = "Please enter your email address.";
        }
        if (!this.form.League) {
            this.errors.League = "Please enter your league's name.";
        }
        if (!this.form.Message) {
            this.errors.Message = "Please enter a message.";
        }
        if (Object.keys(this.errors).length > 0) {
            return;
        }

        // Check that the email field has a valid email address
        const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
        if (!emailRegex.test(this.form.Email)) {
            this.errors.Email = "Please enter a valid email address.";
            return;
        }

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

<style>
    .form-group {
        margin-bottom: 20px;
    }

    .input-wrapper {
        display: flex;
        flex-direction: column;
    }

    input[type="text"],
    input[type="email"],
    textarea {
        display: block;
        width: 95%;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 5px;
        font-size: 16px;
        margin-bottom: 10px;
        outline-color: rgba(255, 159, 33, 0.73);
    }

    textarea {
        height: 150px;
    }

    .error {
        color: red;
        font-size: 14px;
        margin-top: -4px;
        margin-bottom: -12px;
    }

    .footer__success {
        background-color: #dff0d8;
        color: #000000;
        border: 1px solid #000000;
        font-family: "VT323", monospace;
        padding: 20px;
        font-size: large;
        margin-top: 20px;
    }
</style>