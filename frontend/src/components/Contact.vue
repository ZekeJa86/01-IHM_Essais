<template>
<div class="isolate bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900 px-6 py-24 sm:py-10 lg:px-8">
  <div aria-hidden="true" class="absolute inset-x-0 -top-0 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-20">
    <div style="clip-path: polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)" class="relative left-1/2 -z-10 aspect-1155/678 w-144.5 max-w-none -translate-x-1/2 rotate-30 bg-linear-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-40rem)] sm:w-288.75"></div>
  </div>
  <div class="mx-auto max-w-2xl text-center">
    <h2 class="text-4xl font-semibold tracking-tight text-balance text-white-900 sm:text-5xl">Go/No-Go Zone</h2>
    <p class="mt-2 text-lg/8 text-gray-600">Modernize testing, secure the future</p>
  </div>

  <!-- Message de confirmation/erreur -->
  <div v-if="showMessage" :class="messageClass" class="mx-auto mt-4 max-w-xl rounded-md p-4">
    <p class="text-sm font-medium">{{ messageText }}</p>
  </div>

  <form @submit.prevent="handleSubmit" class="mx-auto mt-10 max-w-xl sm:mt-5">
    <div class="grid grid-cols-1 gap-x-8 gap-y-2 sm:grid-cols-2">
      <div>
        <label for="first-name" class="block text-sm/6 font-semibold text-white-900">First name</label>
        <div class="mt-2.5">
          <input 
            v-model="formData.firstName" 
            id="first-name" 
            type="text" 
            name="first-name" 
            autocomplete="given-name" 
            required
            :disabled="isSubmitting"
            class="block w-full rounded-md bg-white px-3.5 py-2 text-base text-black outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 disabled:opacity-50" 
          />
        </div>
      </div>
      <div>
        <label for="last-name" class="block text-sm/6 font-semibold text-white-900">Last name</label>
        <div class="mt-2.5">
          <input 
            v-model="formData.lastName" 
            id="last-name" 
            type="text" 
            name="last-name" 
            autocomplete="family-name" 
            required
            :disabled="isSubmitting"
            class="block w-full rounded-md bg-white px-3.5 py-2 text-base text-black outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 disabled:opacity-50" 
          />
        </div>
      </div>
      <div class="sm:col-span-2">
        <label for="company" class="block text-sm/6 font-semibold text-white-900">Company</label>
        <div class="mt-1">
          <input 
            v-model="formData.company" 
            id="company" 
            type="text" 
            name="company" 
            autocomplete="organization"
            :disabled="isSubmitting"
            class="block w-full rounded-md bg-white px-3.5 py-2 text-base text-black outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 disabled:opacity-50" 
          />
        </div>
      </div>
      <div class="sm:col-span-2">
        <label for="email" class="block text-sm/6 font-semibold text-white-900">Email</label>
        <div class="mt-1">
          <input 
            v-model="formData.email" 
            id="email" 
            type="email" 
            name="email" 
            autocomplete="email" 
            required
            :disabled="isSubmitting"
            class="block w-full rounded-md bg-white px-3.5 py-2 text-base text-black outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 disabled:opacity-50" 
          />
        </div>
      </div>
      <div class="sm:col-span-2">
        <label for="phone-number" class="block text-sm/6 font-semibold text-white-900">Phone number</label>
        <div class="mt-1">
          <div class="flex rounded-md bg-white outline-1 -outline-offset-1 outline-gray-300 has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-600">
            <div class="grid shrink-0 grid-cols-1 focus-within:relative">
              <select 
                v-model="formData.country" 
                id="country" 
                name="country" 
                autocomplete="country" 
                aria-label="Country"
                :disabled="isSubmitting"
                class="col-start-1 row-start-1 w-full appearance-none rounded-md py-2 pr-7 pl-3.5 text-base text-gray-500 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6 disabled:opacity-50"
              >
                <option>EU</option>
                <option>CA</option>
                <option>US</option>
              </select>
              <svg viewBox="0 0 16 16" fill="currentColor" data-slot="icon" aria-hidden="true" class="pointer-events-none col-start-1 row-start-1 mr-2 size-5 self-center justify-self-end text-gray-500 sm:size-4">
                <path d="M4.22 6.22a.75.75 0 0 1 1.06 0L8 8.94l2.72-2.72a.75.75 0 1 1 1.06 1.06l-3.25 3.25a.75.75 0 0 1-1.06 0L4.22 7.28a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>
            </div>
            <input 
              v-model="formData.phoneNumber" 
              id="phone-number" 
              type="text" 
              name="phone-number" 
              placeholder="123-456-7890"
              :disabled="isSubmitting"
              class="block min-w-0 grow py-1.5 pr-3 pl-1 text-base text-black placeholder:text-gray-400 focus:outline-none sm:text-sm/6 disabled:opacity-50" 
            />
          </div>
        </div>
      </div>
      <div class="sm:col-span-2">
        <label for="message" class="block text-sm/6 font-semibold text-white-900">Message</label>
        <div class="mt-1">
          <textarea 
            v-model="formData.message" 
            id="message" 
            name="message" 
            rows="4" 
            required
            :disabled="isSubmitting"
            class="block w-full rounded-md bg-white px-3.5 py-2 text-base text-black outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 disabled:opacity-50"
          ></textarea>
        </div>
      </div>
      <div class="flex gap-x-4 sm:col-span-2">
        <div class="flex h-6 items-center">
          <div class="group relative inline-flex w-8 shrink-0 rounded-full bg-gray-200 p-px inset-ring inset-ring-gray-900/5 outline-offset-2 outline-indigo-600 transition-colors duration-200 ease-in-out has-checked:bg-indigo-600 has-focus-visible:outline-2">
            <span class="size-4 rounded-full bg-white shadow-xs ring-1 ring-gray-900/5 transition-transform duration-200 ease-in-out group-has-checked:translate-x-3.5"></span>
            <input 
              v-model="formData.agreeToPolicy" 
              id="agree-to-policies" 
              type="checkbox" 
              name="agree-to-policies" 
              aria-label="Agree to policies" 
              required
              :disabled="isSubmitting"
              class="absolute inset-0 appearance-none focus:outline-hidden" 
            />
          </div>
        </div>
        <label for="agree-to-policies" class="text-sm/6 text-gray-600">
          By selecting this, you agree to our
          <a href="#" class="font-semibold whitespace-nowrap text-indigo-600">privacy policy</a>.
        </label>
      </div>
    </div>
    <div class="mt-10">
      <button 
        type="submit" 
        :disabled="isSubmitting"
        class="block w-full rounded-md bg-indigo-600 px-3.5 py-2.5 text-center text-sm font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
      >
        <span v-if="!isSubmitting">Let's talk</span>
        <span v-else>Envoi en cours...</span>
      </button>
    </div>
  </form>
</div>
</template>

<script setup>
import { ref } from 'vue';
import DarkMode from './DarkMode.vue';
import Header from './Header.vue';

console.log("Contact.vue monté ✅");

// Données du formulaire
const formData = ref({
  firstName: '',
  lastName: '',
  company: '',
  email: '',
  country: 'EU',
  phoneNumber: '',
  message: '',
  agreeToPolicy: false
});

// État des messages et du formulaire
const showMessage = ref(false);
const messageText = ref('');
const messageClass = ref('');
const isSubmitting = ref(false);

// Fonction pour afficher un message
const displayMessage = (text, type) => {
  messageText.value = text;
  messageClass.value = type === 'success' 
    ? 'bg-green-50 text-green-800 border border-green-200' 
    : 'bg-red-50 text-red-800 border border-red-200';
  showMessage.value = true;
  
  setTimeout(() => {
    showMessage.value = false;
  }, 5000);
};

// Fonction de soumission du formulaire
const handleSubmit = async () => {
  // Vérification de la case à cocher
  if (!formData.value.agreeToPolicy) {
    displayMessage('Veuillez accepter notre politique de confidentialité', 'error');
    return;
  }

  isSubmitting.value = true;

  try {
    // Vérifier que l'objet window.go existe (Wails)
    if (!window.go || !window.go.main || !window.go.main.App) {
      throw new Error('Wails runtime non disponible. Êtes-vous en mode dev ?');
    }

    // Appel à la fonction Go via Wails
    const result = await window.go.main.App.SendEmail(formData.value);

    if (result.success) {
      displayMessage('✅ ' + result.message, 'success');
      
      // Réinitialisation du formulaire
      formData.value = {
        firstName: '',
        lastName: '',
        company: '',
        email: '',
        country: 'EU',
        phoneNumber: '',
        message: '',
        agreeToPolicy: false
      };
    } else {
      displayMessage('❌ ' + result.message, 'error');
    }
  } catch (error) {
    console.error('Erreur:', error);
    displayMessage('❌ Erreur lors de l\'envoi: ' + error.message, 'error');
  } finally {
    isSubmitting.value = false;
  }
};
</script>