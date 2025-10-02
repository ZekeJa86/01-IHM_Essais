<template>
  <div class="h-full flex flex-col">
    <!-- Section IMAGES -->
    <div class="flex-[6] bg-slate-900 flex items-center justify-center text-white p-4">
      <div class="flex justify-between w-full">
        <div class="bg-slate-700 flex items-center justify-center w-[30%] h-40">
          <img :src="usineImg" class="w-[30%] h-40 object-cover" alt="Image 1" />
        </div>
        <div class="bg-slate-700 flex items-center justify-center w-[30%] h-40">
          <img :src="usineImg" class="w-[30%] object-cover" alt="Image 1" />
        </div>
        <div class="bg-slate-700 flex items-center justify-center w-[30%] h-40">
         <img :src="usineImg" class="w-[30%] h-40 object-cover" alt="Image 1" />
        </div>
      </div>
    </div>

    <!-- Section EPI -->
    <div class="flex-[0.2] bg-slate-600 flex items-center justify-center text-white py-8">
      EPI
    </div>

    <!-- Section CONSIGNES -->
    <div class="flex-[2] bg-slate-900 flex items-center justify-center text-white">
      CONSIGNES
    </div>
  </div>
</template>



<script setup>
import usineImg from '../components/Images/usine.jpg'
</script>

