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

            <img 
              :src="myselfImage" 
              alt="Photo de profil KABRÉ" 
              class="w-full h-full object-cover"
            />
            <!-- Overlay avec texte au survol -->
            <div class="absolute inset-0 flex flex-col items-center justify-end text-center px-6 pb-12 bg-slate-800/50">
              <h1 class="text-5xl md:text-6xl font-bold text-white mb-4 tracking-wider animate-fade-in">
                KABRÉ Sylvain
              </h1>
              
              <p class="text-lg md:text-2xl text-gray-200 leading-relaxed max-w-3xl animate-fade-in-delay">
                « Ensemble, donnons vie à vos projets <a class="text-blue-400 hover:text-blue-300 transition-colors font-semibold" href="#">numériques</a>. » 
              </p>

            </div>
          </div>
        </div>

        <!-- Boutons d'action sous l'image -->
        <div class="flex flex-wrap gap-4 justify-center py-12 animate-fade-in-delay-2">
          <a href="/contact" class="px-8 py-3 bg-gradient-to-r from-blue-500 to-purple-500 text-white font-semibold rounded-lg hover:shadow-lg hover:shadow-purple-500/50 transition-all duration-300">
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

    <!-- Statistiques réelles et vendeuses -->
    <div class="relative py-20 bg-gray-900">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-3xl font-bold text-center mb-12 text-white scroll-reveal">Mon expertise en chiffres</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
          <div class="text-center group scroll-reveal" style="--delay: 0.1s">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedYears }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Ans d'ingénierie</div>
          </div>
          <div class="text-center group scroll-reveal" style="--delay: 0.2s">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedTech }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Technologies maîtrisées</div>
          </div>
          <div class="text-center group scroll-reveal" style="--delay: 0.3s">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              {{ animatedProject }}+
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Projets menés</div>
          </div>
          <div class="text-center group scroll-reveal" style="--delay: 0.4s">
            <div class="text-5xl font-bold text-blue-400 mb-2 transition-transform duration-300 group-hover:scale-110">
              100<span class="text-2xl">%</span>
            </div>
            <div class="text-gray-400 text-sm uppercase tracking-wider">Engagement qualité</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Section Valeur ajoutée -->
    <div class="relative py-20 bg-gray-800">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white scroll-reveal">Pourquoi me choisir ?</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12 scroll-reveal"></div>

        <div class="grid md:grid-cols-3 gap-6">
          <!-- Avantage 1 -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.1s">
            <div class="text-4xl mb-4">🎯</div>
            <h3 class="text-xl font-semibold text-white mb-4">Double expertise</h3>
            <p class="text-gray-400 text-sm leading-relaxed">
              Ingénieur de formation avec 5 ans d'expérience industrielle, je combine rigueur technique et créativité pour des solutions sur-mesure.
            </p>
          </div>

          <!-- Avantage 2 -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.2s">
            <div class="text-4xl mb-4">⚡</div>
            <h3 class="text-xl font-semibold text-white mb-4">Technologies modernes</h3>
            <p class="text-gray-400 text-sm leading-relaxed">
              Stack technique à jour (Vue.js, React, Go, Tailwind) pour des applications performantes et évolutives.
            </p>
          </div>

          <!-- Avantage 3 -->
          <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.3s">
            <div class="text-4xl mb-4">🤝</div>
            <h3 class="text-xl font-semibold text-white mb-4">Accompagnement complet</h3>
            <p class="text-gray-400 text-sm leading-relaxed">
              De l'idée à la mise en production, je vous accompagne à chaque étape avec transparence et réactivité.
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Section Compétences techniques -->
    <div class="relative py-20 bg-gray-900">
      <div class="max-w-6xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white scroll-reveal">Stack Technique</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12 scroll-reveal"></div>

        <div class="grid md:grid-cols-3 gap-6">
          <!-- Frontend -->
          <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.1s">
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
                <span class="text-gray-300">EmailJS</span>
              </div>
            </div>
          </div>

          <!-- Backend -->
          <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.2s">
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
                  <span class="text-blue-400 font-bold">N</span>
                </div>
                <span class="text-gray-300">Node.js</span>
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
          <div class="bg-gray-800/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 scroll-reveal" style="--delay: 0.3s">
            <div class="text-3xl mb-4">🚀</div>
            <h3 class="text-xl font-semibold text-white mb-4">DevOps & Tools</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-500/20 rounded-lg flex items-center justify-center">
                  <span class="text-blue-400 font-bold">G</span>
                </div>
                <span class="text-gray-300">Git / GitHub</span>
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
    <div id="parcours" class="relative py-20 bg-gray-800">
      <div class="max-w-7xl mx-auto px-4">
        <h2 class="text-4xl font-bold text-center mb-4 text-white scroll-reveal">Mon Parcours</h2>
        <div class="h-1 w-32 bg-gradient-to-r from-blue-500 to-blue-400 mx-auto mb-12 scroll-reveal"></div>

        <div class="grid md:grid-cols-2 gap-8">
          <!-- Colonne Expérience Professionnelle -->
          <div class="relative">
            <div class="sticky top-8 mb-8 scroll-reveal">
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
              <div class="relative pl-8 scroll-reveal" style="--delay: 0.1s">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-800"></div>
                <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Ingénieur développement solutions de test</h4>
                  <p class="text-blue-400 mb-3 font-semibold">2023 - Présent</p>
                  <p class="text-gray-400 text-sm leading-relaxed">Développement d'applications de test pour l'industrie de la batterie. Automatisation et optimisation des processus.</p>
                </div>
              </div>

              <!-- Expérience 2 -->
              <div class="relative pl-8 scroll-reveal" style="--delay: 0.2s">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-blue-500 rounded-full border-4 border-gray-800"></div>
                <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Ingénieur bureau d'études</h4>
                  <p class="text-blue-400 mb-3 font-semibold">2019 - 2022</p>
                  <p class="text-gray-400 text-sm leading-relaxed">Conception électromécanique dans le secteur des ascenseurs. Gestion de projets techniques complexes.</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Colonne KDEV -->
          <div class="relative">
            <div class="sticky top-8 mb-8 scroll-reveal">
              <h3 class="text-2xl font-bold text-purple-400 mb-2 flex items-center gap-3">
                <span class="text-3xl">👨🏽‍💼</span>
                KDEV - Freelance
              </h3>
              <div class="h-1 w-20 bg-purple-500"></div>
            </div>

            <div class="relative space-y-8">
              <!-- Ligne verticale -->
              <div class="absolute left-0 top-0 bottom-0 w-0.5 bg-purple-500/30"></div>

              <!-- Activité actuelle -->
              <div class="relative pl-8 scroll-reveal" style="--delay: 0.1s">
                <div class="absolute left-0 transform -translate-x-1/2 w-4 h-4 bg-purple-500 rounded-full border-4 border-gray-800"></div>
                <div class="bg-gray-900/50 backdrop-blur-sm border border-gray-700 rounded-xl p-6 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1">
                  <h4 class="text-xl font-semibold text-white mb-2">Développeur Full-Stack</h4>
                  <p class="text-purple-400 mb-3 font-semibold">2025 - Présent</p>
                  <p class="text-gray-400 text-sm leading-relaxed mb-2">
                    <strong class="text-white">Applications web modernes</strong> : Sites vitrines, applications métier, e-commerce
                  </p>
                  <p class="text-gray-400 text-sm leading-relaxed">
                    <strong class="text-white">Modernisation</strong> : Refonte et optimisation d'applications de bureau pour industriels
                  </p>
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
import myselfImage from '../assets/images/myself.png'

export default {
  name: 'ProfilePage',
  setup() {
    const animatedYears = ref(0);
    const animatedTech = ref(0);
    const animatedProject = ref(0);

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

    // Fonction pour gérer les animations au scroll
    const setupScrollAnimations = () => {
      const scrollRevealElements = document.querySelectorAll('.scroll-reveal');
      
      const scrollObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            entry.target.classList.add('revealed');
          }
        });
      }, {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
      });

      scrollRevealElements.forEach(el => {
        scrollObserver.observe(el);
      });
    };

    onMounted(() => {
      // Animation des statistiques
      const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            animateValue(animatedYears, 5);
            animateValue(animatedTech, 15);
            animateValue(animatedProject, 50);
            observer.disconnect();
          }
        });
      });

      const statsElement = document.querySelector('.grid.grid-cols-2');
      if (statsElement) {
        observer.observe(statsElement);
      }

      // Setup des animations au scroll
      setupScrollAnimations();
    });

    return {
      myselfImage,
      animatedYears,
      animatedTech,
      animatedProject
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

/* Animations Hero */
@keyframes fadeIn {
  from { 
    opacity: 0; 
    transform: translateY(20px);
  }
  to { 
    opacity: 1; 
    transform: translateY(0);
  }
}

.animate-fade-in {
  animation: fadeIn 1s ease-out;
}

.animate-fade-in-delay {
  animation: fadeIn 1s ease-out 0.3s backwards;
}

.animate-fade-in-delay-2 {
  animation: fadeIn 1s ease-out 0.6s backwards;
}

/* Animations au scroll */
.scroll-reveal {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.6s ease-out, transform 0.6s ease-out;
  transition-delay: var(--delay, 0s);
}

.scroll-reveal.revealed {
  opacity: 1;
  transform: translateY(0);
}

.animate-pulse-slow {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

html {
  scroll-behavior: smooth;
}
</style>