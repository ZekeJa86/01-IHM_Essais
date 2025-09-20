<script setup>
import { useDark } from '@vueuse/core'

// Liaison de isDark avec la classe dark sur <html>
const isDark = useDark(document.documentElement, false)
</script>


<template>
  <div class="min-h-screen flex items-center justify-center bg-white dark:bg-gray-900 transition-colors duration-500">
    <!-- Switch animé -->
    <div class="switch">
      <input type="checkbox" id="Switch" class="switch__input" v-model="isDark" />
      <label for="Switch" class="switch__label">
        <span class="switch__indicator"></span>
        <span class="switch__decoration"></span>
      </label>
    </div>
  </div>
</template>




<style>
.body {
  padding: 25px; /* moitié de 50px */
}

.switch {
  display: inline-block;
  position: relative;
}

.switch__input {
  clip: rect(1px, 1px, 1px, 1px);
  clip-path: inset(50%);
  height: 1px;
  width: 1px;
  margin: -1px;
  overflow: hidden;
  padding: 0;
  position: absolute;
}

.switch__label {
  position: relative;
  display: inline-block;
  width: 60px;  /* 120 / 2 */
  height: 30px; /* 60 / 2 */
  background-color: #2B2B2B;
  border: 2.5px solid #5B5B5B; /* 5 / 2 */
  border-radius: 9999px;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(.46,.03,.52,.96);
}

.switch__indicator {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) translateX(-72%); /* garder ratio original */
  width: 20px;  /* 40 / 2 */
  height: 20px; /* 40 / 2 */
  background-color: #7B7B7B;
  border-radius: 9999px;
  box-shadow: 5px 0 0 0 rgba(0,0,0,0.2) inset; /* 10 / 2 */
  transition: all 0.4s cubic-bezier(.46,.03,.52,.96);
}

.switch__indicator::before,
.switch__indicator::after {
  position: absolute;
  content: '';
  display: block;
  background-color: #FFFFFF;
  border-radius: 9999px;
  transition: all 0.4s cubic-bezier(.46,.03,.52,.96);
}

.switch__indicator::before {
  top: 3.5px;   /* 7 / 2 */
  left: 3.5px;  /* 7 / 2 */
  width: 4.5px; /* 9 / 2 */
  height: 4.5px;
  opacity: 0.6;
}

.switch__indicator::after {
  bottom: 4px; /* 8 / 2 */
  right: 3px;  /* 6 / 2 */
  width: 7px;  /* 14 / 2 */
  height: 7px;
  opacity: 0.8;
}

.switch__decoration {
  position: absolute;
  top: 32.5%; /* 65 / 2 */
  left: 50%;
  width: 2.5px;  /* 5 / 2 */
  height: 2.5px; /* 5 / 2 */
  background-color: #FFFFFF;
  border-radius: 9999px;
  animation: twinkle 0.8s infinite -0.6s;
  transition: all 0.4s cubic-bezier(.46,.03,.52,.96);
}

.switch__decoration::before,
.switch__decoration::after {
  position: absolute;
  display: block;
  content: '';
  width: 2.5px;  /* 5 / 2 */
  height: 2.5px;
  background-color: #FFFFFF;
  border-radius: 9999px;
}

.switch__decoration::before {
  top: -10px;  /* -20 / 2 */
  left: 5px;   /* 10 / 2 */
  opacity: 1;
  animation: twinkle 0.6s infinite;
}

.switch__decoration::after {
  top: -3.5px; /* -7 / 2 */
  left: 15px;  /* 30 / 2 */
  animation: twinkle 0.6s infinite -0.2s;
}

@keyframes twinkle {
  50% { opacity: 0.2; }
}

.switch__input:checked + .switch__label {
  background-color: #8FB5F5;
  border-color: #347CF8;
}

.switch__input:checked + .switch__label .switch__indicator {
  background-color: #ECD21F;
  box-shadow: none;
  transform: translate(-50%, -50%) translateX(72%); /* garder ratio original */
}

.switch__input:checked + .switch__label .switch__indicator::before,
.switch__input:checked + .switch__label .switch__indicator::after {
  display: none;
}

.switch__input:checked + .switch__label .switch__decoration {
  top: 25%; /* 50 / 2 */
  transform: translate(0%, -50%);
  animation: cloud 8s linear infinite;
  width: 10px;  /* 20 / 2 */
  height: 10px; /* 20 / 2 */
}

.switch__input:checked + .switch__label .switch__decoration::before {
  width: 5px;  /* 10 / 2 */
  height: 5px;
  top: auto;
  bottom: 0;
  left: -4px; /* -8 / 2 */
  animation: none;
}

.switch__input:checked + .switch__label .switch__decoration::after {
  width: 7.5px;  /* 15 / 2 */
  height: 7.5px;
  top: auto;
  bottom: 0;
  left: 8px;  /* 16 / 2 */
  border-bottom-right-radius: 9999px;
}

@keyframes cloud {
  0% { transform: translate(0%, -50%); }
  50% { transform: translate(-50%, -50%); }
  100% { transform: translate(0%, -50%); }
}

</style>