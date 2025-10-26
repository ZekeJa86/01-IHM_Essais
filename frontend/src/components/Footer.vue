<template>
<footer class="flex flex-row dark:bg-gray-900 py-1 w-full">
    <div class="w-full max-w-screen-xl mx-auto px-4">
        <div class="sm:flex sm:items-center sm:justify-between">
            <a class="flex items-center mb-4 sm:mb-0 space-x-3 rtl:space-x-reverse">
                <img src="../assets/images/KDEV_light.png" alt="Logo KDEV" class="h-12 w-auto" />
                <span class="self-center text-sm font-semibold whitespace-nowrap dark:text-white">« Ensemble, donnons vie à vos projets numériques. »</span>
            </a>
            <ul class="flex flex-wrap items-center text-sm font-medium text-gray-500 dark:text-gray-400">
                <li>
                    <a href="#" class="hover:underline me-4 md:me-6">
                        <router-link to="/portfolio" :class="navLinkClass">Portfolio</router-link>
                    </a>
                </li>
                <li>
                    <a href="#" class="hover:underline me-4 md:me-6">
                         <router-link to="/privacy" :class="navLinkClass">Privacy Policy</router-link>
                    </a>
                </li>
                <li>
                    <a href="#" class="hover:underline">
                         <router-link to="/contact" :class="navLinkClass">Contact</router-link>
                    </a>
                </li>
            </ul>
        </div>
        <hr class="my-2 border-gray-200 sm:mx-auto dark:border-gray-700" />
        <span class="block text-sm text-gray-500 sm:text-center dark:text-gray-400 pb-2">© 2025 <a class="hover:underline">KDEV</a>. All Rights Reserved.</span>
    </div>
</footer>

</template>

<script setup>
console.log("Footer.vue monté ✅");
</script>