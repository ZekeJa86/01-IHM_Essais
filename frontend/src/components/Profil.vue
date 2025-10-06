<template>
  <div class="h-full flex flex-col">
    <!-- Section IMAGES -->
    <div class="pt-30 flex-1 flex flex-col items-center justify-center text-white p-4 space-y-8">
      
      <!-- Card principale -->
      <div class="card group hover:shadow-sm sm:max-w-3xl w-full">
        <figure class="relative max-w-sm mx-auto transition-all duration-300 cursor-pointer filter grayscale hover:grayscale-0 flex justify-center">
          <img src="../components/Images/myself.jpeg" alt="Profile" class="rounded-lg transition-transform duration-500 h-100 group-hover:scale-110 mx-auto" />
        </figure>
        <div class="card-body">
          <!-- Titre manuscrit centré -->
          <h5 class="tracking-wide card-title manrope mb-2 pt-10 text-3xl mx-auto text-center">
            KABRÉ SYLVAIN
          </h5>

          <p class="text-xl pt-5">
            Ma passion pour la <a class="text-blue-600 after:content-['']" href="...">
              technologie</a> va de pair avec ma recherche constante de défis intellectuels.
            Mon parcours <a class="text-blue-600 after:content-['']" href="...">professionnel</a>
            s'est construit autour de la résolution de problèmes complexes, une démarche qui a façonné mes compétences. 
            Si vous partagez ces aspirations ou souhaitez échanger sur des sujets technologiques, n'hésitez pas à me 
            <a class="text-blue-600 after:content-['']" href="...">contacter.</a>
          </p>
        </div>
      </div>
      </div>
    </div>
</template>


<style>
.bg_photo1 {
  background-image: url('../components/Images/photo1.png');
  background-size: cover;
  background-position: center;
}

.bg_photo2 {
  background-image: url('../components/Images/photo2.png');
  background-size: cover;
  background-position: center;
}

.bg_photo3 {
  background-image: url('../components/Images/photo3.png');
  background-size: cover;
  background-position: center;
}
</style>