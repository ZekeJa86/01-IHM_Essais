<template>
  <div class="h-full flex flex-col">
    <!-- Section IMAGES -->
    <div class="pt-10 flex-[6] bg-slate-900 flex items-center justify-center text-white p-4">
      <div class="card group hover:shadow-sm sm:max-w-sm">
        <figure class="relative max-w-sm transition-all duration-300 cursor-pointer filter grayscale hover:grayscale-0">
          <img src="../components/Images/myself.jpeg" alt="Shoes" class="rounded-lg transition-transform duration-500 group-hover:scale-110" />
        </figure>
        <div class="card-body">
          <h5 class="card-title mb-2.5 pt-5">Card title</h5>
          <p class="mb-6">Ma passion pour la technologie va de pair avec ma recherche constante de défis intellectuels. Mon parcours professionnel s'est construit autour de la résolution de problèmes complexes, une démarche qui a façonné mes compétences. Si vous partagez ces aspirations ou souhaitez échanger sur des sujets technologiques, n'hésitez pas à me contacter.</p>
          <div class="card-actions">
            <button class="btn btn-primary">Buy Now</button>
            <button class="btn btn-soft">Add to cart</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Footer />
</template>

