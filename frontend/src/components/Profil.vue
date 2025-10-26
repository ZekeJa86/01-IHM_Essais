<template>
  <div>
    <!-- Hero Section avec profil -->
    <div class="relative min-h-screen flex flex-col bg-gradient-to-b from-gray-900 to-gray-800 overflow-hidden">
      <!-- Grille animée de fond -->
      <div class="absolute inset-0 opacity-10">
        <div class="absolute inset-0 grid-background"></div>
      </div>

      <!-- Orbes lumineux -->
      <div class="absolute top-20 left-20 w-96 h-96 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full filter blur-3xl opacity-20 animate-pulse"></div>
      <div class="absolute bottom-20 right-20 w-96 h-96 bg-gradient-to-r from-purple-400 to-blue-600 rounded-full filter blur-3xl opacity-15 animate-pulse-slow"></div>

      <!-- Section Hero simple -->
      <div class="relative z-10 flex-1 flex flex-col items-center justify-center text-white">
        <!-- Image pleine largeur -->
        <div class="relative w-full mx-auto group">
          <div class="relative w-full h-[60vh] md:h-[70vh] overflow-hidden not-[]:shadow-lg">

            <img :src="myselfImage" />


            <!-- Overlay avec texte au survol -->
            <div class="absolute inset-0 flex flex-col items-center justify-end text-center px-6 pb-12 bg-slate-800/50">
              <h1 class="text-5xl md:text-6xl font-bold text-white mb-4 tracking-wider">
                KABRÉ
              </h1>
              
              <p class="text-lg md:text-xl text-gray-200 leading-relaxed max-w-3xl">
                Ma passion pour la <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold" href="#">technologie</a> va de pair avec ma recherche constante de défis intellectuels.
                Mon parcours <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold" href="#">professionnel</a> s'est construit autour de la résolution de problèmes complexes.
                N'hésitez pas à me <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold" href="#contact">contacter</a>.
              </p>
            </div>
          </div>
        </div>

        <!-- Boutons d'action sous l'image -->
        <div class="flex flex-wrap gap-4 justify-center py-12">
          <a href="#contact" class="px-8 py-3 bg-gradient-to-r from-blue-500 to-purple-500 text-white font-semibold rounded-lg hover:shadow-lg hover:shadow-purple-500/50 transition-all duration-300">
            Me contacter
          </a>
          <a href="#parcours" class="px-8 py-3 border-2 border-blue-400 text-blue-400 font-semibold rounded-lg hover:bg-blue-400 hover:text-white transition-all duration-300">
            Voir mon parcours
          </a>
        </div>
      </div>

      <!-- Indicateur de scroll -->
      <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce">
        <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
      </div>
    </div>

    <!-- Statistiques animées -->
    <div class="relative py-20 bg-gray-900">
      <div class="max-w-6xl mx-auto px-4">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedYears }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Années d'expérience</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedProjects }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Projets réalisés</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedTech }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Technologies</div>
          </div>
          <div class="text-center group">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedClients }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Clients satisfaits</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Section Compétences -->
    <div class="relative py-20 bg-gray-800">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white">Expertise Technique</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12"></div>

        <div class="grid md:grid-cols-3 gap-6">
          <!-- Frontend -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">⚡</div>
            <h3 class="text-xl font-semibold text-white mb-4">Frontend</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold text-xs">VJS</span>
                </div>
                <span class="text-gray-300">Vue.js</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">R</span>
                </div>
                <span class="text-gray-300">React</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">T</span>
                </div>
                <span class="text-gray-300">Tailwind CSS</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold text-xs">EJS</span>
                </div>
                <span class="text-gray-300">Email JS</span>
              </div>
            </div>
          </div>

          <!-- Backend -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">🔧</div>
            <h3 class="text-xl font-semibold text-white mb-4">Backend</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">Go</span>
                </div>
                <span class="text-gray-300">Golang</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">V</span>
                </div>
                <span class="text-gray-300">Vite</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">D</span>
                </div>
                <span class="text-gray-300">Databases</span>
              </div>
            </div>
          </div>

          <!-- DevOps -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2">
            <div class="text-3xl mb-4">🚀</div>
            <h3 class="text-xl font-semibold text-white mb-4">DevOps & Tools</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">G</span>
                </div>
                <span class="text-gray-300">Github</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold text-xs">VS</span>
                </div>
                <span class="text-gray-300">VS Code</span>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">V</span>
                </div>
                <span class="text-gray-300">Vercel</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Timeline Parcours -->
    <div id="parcours" class="relative py-20 bg-gray-900">
      <div class="max-w-7xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white">Parcours</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12"></div>

        <div class="grid md:grid-cols-2 gap-8">
          <!-- Colonne Expérience Professionnelle -->
          <div class="relative">
            <div class="sticky top-8 mb-8">
              <h3 class="text-2xl font-bold text-blue-400 mb-2 flex items-center gap-3">
                <span class="text-3xl">💼</span>
                Expérience Professionnelle
              </h3>
              <div class="h-1 w-20 bg-blue-500"></div>
            </div>

            <div class="relative space-y-8">
              <!-- Ligne verticale -->
              <div class="absolute left-0 top-0 bottom-0 w-0.5 bg-blue-500/30"></div>

              <!-- Expérience 1 -->
              <div class="relative pl-8">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-900"></div>
                <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Ingénieur en développement de solutions de test</h4>
                  <p class="text-blue-400 mb-3 font-semibold">2023 - Présent</p>
                  <p class="text-gray-400 text-sm leading-relaxed">Développement d'applications de test pour l'industrie automobile</p>
                </div>
              </div>

              <!-- Expérience 2 -->
              <div class="relative pl-8">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-900"></div>
                <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Ingénieur bureau d'études</h4>
                  <p class="text-blue-400 mb-3 font-semibold">2019 - 2022</p>
                  <p class="text-gray-400 text-sm leading-relaxed">Électromécaniques dans le milieu des ascensoristes</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Colonne KDEV -->
          <div class="relative">
            <div class="sticky top-8 mb-8">
              <h3 class="text-2xl font-bold text-purple-400 mb-2 flex items-center gap-3">
                <span class="text-3xl">👨🏽‍💼</span>
                KDEV
              </h3>
              <div class="h-1 w-20 bg-purple-500"></div>
            </div>

            <div class="relative space-y-8">
              <!-- Ligne verticale -->
              <div class="absolute left-0 top-0 bottom-0 w-0.5 bg-purple-500/30"></div>

              <!-- Formation 1 -->
              <div class="relative pl-8">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-purple-500 rounded-full border-4 border-gray-900"></div>
                <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Développeur Full-Stack</h4>
                  <p class="text-purple-400 mb-3 font-semibold">2025 - Présent</p>
                  <p class="text-gray-400 text-sm leading-relaxed">Création d'applications web modernes et performantes</p>
                  <p class="text-gray-400 text-sm leading-relaxed mt-2">Création et modernisation d'applications de bureau pour l'industrie</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import myselfImage from '../assets/Images/myself.png'

export default {
  name: 'ProfilePage',
  setup() {
    const animatedYears = ref(0);
    const animatedProjects = ref(0);
    const animatedTech = ref(0);
    const animatedClients = ref(0);

    const animateValue = (ref, target, duration = 2000) => {
      const start = 0;
      const increment = target / (duration / 16);
      let current = start;

      const timer = setInterval(() => {
        current += increment;
        if (current >= target) {
          ref.value = target;
          clearInterval(timer);
        } else {
          ref.value = Math.floor(current);
        }
      }, 16);
    };

    onMounted(() => {
      const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            animateValue(animatedYears, 5);
            animateValue(animatedProjects, 50);
            animateValue(animatedTech, 20);
            animateValue(animatedClients, 30);
            observer.disconnect();
          }
        });
      });

      const statsElement = document.querySelector('.grid.grid-cols-2');
      if (statsElement) {
        observer.observe(statsElement);
      }
    });

    return {
      animatedYears,
      animatedProjects,
      animatedTech,
      animatedClients
    };
  }
};
</script>

<style scoped>
.grid-background {
  background-image: linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
                    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
  animation: gridMove 20s linear infinite;
}

@keyframes gridMove {
  0% { transform: translateY(0); }
  100% { transform: translateY(50px); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.animate-pulse-slow {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

html {
  scroll-behavior: smooth;
}
</style>