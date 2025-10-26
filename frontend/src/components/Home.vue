<template>
  <div class="relative bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900 w-full px-4 overflow-hidden">
    <!-- Animated Background -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <!-- Gradient orbs -->
      <div class="absolute top-1/4 -left-48 w-96 h-96 bg-blue-500/10 rounded-full blur-3xl animate-pulse"></div>
      <div class="absolute top-1/3 -right-48 w-96 h-96 bg-indigo-500/10 rounded-full blur-3xl animate-pulse" style="animation-delay: 1s;"></div>
      <div class="absolute bottom-1/4 left-1/3 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl animate-pulse" style="animation-delay: 2s;"></div>
      
      <!-- Floating particles -->
      <div class="absolute top-20 left-1/4 w-2 h-2 bg-blue-400/30 rounded-full animate-float"></div>
      <div class="absolute top-40 right-1/3 w-1 h-1 bg-blue-300/40 rounded-full animate-float" style="animation-delay: 0.5s; animation-duration: 4s;"></div>
      <div class="absolute top-60 left-1/2 w-1.5 h-1.5 bg-indigo-400/30 rounded-full animate-float" style="animation-delay: 1s; animation-duration: 5s;"></div>
      <div class="absolute bottom-40 right-1/4 w-2 h-2 bg-cyan-400/30 rounded-full animate-float" style="animation-delay: 1.5s; animation-duration: 6s;"></div>
      <div class="absolute bottom-60 left-1/3 w-1 h-1 bg-blue-300/40 rounded-full animate-float" style="animation-delay: 2s; animation-duration: 4.5s;"></div>
    </div>

    <!-- Hero Section Centered -->
    <div class="min-h-screen flex flex-col items-center justify-center relative z-10">
      <!-- Hero Title -->
      <div
        class="font-extrabold text-5xl md:text-6xl antialiased text-center
              [text-wrap:balance] bg-clip-text text-transparent
              bg-gradient-to-r from-slate-400/60 to-slate-200 mb-6 animate-fade-in">
 Ensemble, donnons vie à vos projets <span class="text-blue-400 inline-flex flex-col">
          <ul class="block text-center leading-tight [&_li]:block">
            <li>numériques</li>
          </ul>
        </span>
      </div>

      <!-- Subtitle -->
      <p class="text-slate-400 text-xl md:text-2xl text-center max-w-3xl mb-12 font-light animate-fade-in-delay">
        Développement web et logiciel sur mesure pour vos environnements industriels.</p>
      <!-- CTA Button -->
      <a href="/contact" 
         class="bg-gradient-to-r from-blue-500 to-purple-500 hover:bg-blue-600 text-white font-semibold px-8 py-4 rounded-full 
                transition-all duration-300 hover:scale-105 hover:shadow-lg hover:shadow-blue-500/50 animate-fade-in-delay-2">
        Demander un devis
      </a>

      <div class="absolute bottom-2 left-1/2 transform -translate-x-1/2 animate-bounce">
        <svg class="w-6 h-6 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
      </div>
    </div>

    <!-- Services Section -->
    <div class="max-w-6xl w-full mx-auto py-20 relative z-10">
      <h2 class="text-4xl font-bold text-center text-slate-200 mb-16 scroll-reveal">Mes Services</h2>
      <h3 class="font-extrabold text-2xl md:text-3xl antialiased text-center
              [text-wrap:balance] bg-clip-text text-transparent
              bg-gradient-to-r from-slate-400/60 to-slate-200 mb-6 animate-fade-in">Solutions dédiées aux environnements
<span class="text-blue-500 inline-flex flex-col">
          <ul class="block text-center leading-tight [&_li]:block">
            <li>industriels</li>
          </ul>
        </span>
          </h3>      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.1s">
          <h3 class="text-blue-400 text-xl font-semibold mb-3">Applications de test</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Développement d'applications de bureau robustes pour automatiser vos processus de test en production
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.2s">
          <h3 class="text-blue-400 text-xl font-semibold mb-3">Intégration système</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Connexion transparente avec vos environnements de production et outils existants
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.3s">
          <h3 class="text-blue-400 text-xl font-semibold mb-3">Monitoring temps réel</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Surveillance continue et rapports détaillés de vos essais en production
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.4s">
          <h3 class="text-blue-400 text-xl font-semibold mb-3">Support & maintenance</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Accompagnement technique et évolutions pour garantir la pérennité de vos solutions
          </p>
        </div>
      </div>
    </div>


<div class="max-w-6xl w-full mx-auto py-20 relative z-10">
      <h3 class="font-extrabold text-2xl md:text-3xl antialiased text-center
              [text-wrap:balance] bg-clip-text text-transparent
              bg-gradient-to-r from-slate-400/60 to-slate-200 mb-6 animate-fade-in">Solutions dédiées aux environnements
            <span class="text-purple-500 inline-flex flex-col">
          <ul class="block text-center leading-tight [&_li]:block">
            <li>web</li>
          </ul>
        </span>
          </h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.1s">
          <h3 class="text-purple-500 text-xl font-semibold mb-3">Sites vitrine & portfolios</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Création de sites web modernes et responsives pour valoriser votre activité
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.2s">
          <h3 class="text-purple-400 text-xl font-semibold mb-3">Applications web SPA</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Développement d'applications single-page avec Vue.js pour des interfaces fluides
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.3s">
          <h3 class="text-purple-400 text-xl font-semibold mb-3">API & Backend</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Conception d'APIs RESTful et services backend performants et sécurisés
          </p>
        </div>

        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-6 
                    hover:bg-slate-800/70 hover:border-purple-500/50 transition-all duration-300 hover:-translate-y-1 scroll-reveal" style="--delay: 0.4s">
          <h3 class="text-purple-400 text-xl font-semibold mb-3">Design UI/UX</h3>
          <p class="text-slate-400 text-sm leading-relaxed">
            Interfaces utilisateur modernes avec Tailwind CSS pour une expérience optimale
          </p>
        </div>
      </div>
    </div>
    

    <!-- Stats Section -->
    <div class="max-w-5xl w-full mx-auto mt-20">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
        <div class="p-4 scroll-reveal" style="--delay: 0.1s">
          <div class="text-5xl font-bold text-blue-400 mb-2">5+</div>
          <div class="text-slate-400 text-sm">Années d'expérience</div>
        </div>
        <div class="p-4 scroll-reveal" style="--delay: 0.2s">
          <div class="text-5xl font-bold text-blue-400 mb-2">15+</div>
          <div class="text-slate-400 text-sm">Technologies maîtrisées</div>
        </div>
        <div class="p-4 scroll-reveal" style="--delay: 0.3s">
          <div class="text-5xl font-bold text-blue-400 mb-2">100%</div>
          <div class="text-slate-400 text-sm">Engagement qualité</div>
        </div>
        <div class="p-4 scroll-reveal" style="--delay: 0.4s">
          <div class="text-5xl font-bold text-blue-400 mb-2">24/7</div>
          <div class="text-slate-400 text-sm">Disponible</div>
        </div>
      </div>
    </div>

    <!-- Why Choose Us Section -->
    <div class="max-w-6xl w-full mx-auto mt-24">
      <h2 class="text-4xl font-bold text-center text-slate-200 mb-16 scroll-reveal">Pourquoi nous choisir ?</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="flex gap-4 scroll-reveal" style="--delay: 0.1s">
          <div class="flex-shrink-0 w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center text-2xl">
            ✓
          </div>
          <div>
            <h3 class="text-slate-200 text-xl font-semibold mb-2">Expertise métier</h3>
            <p class="text-slate-400 text-sm leading-relaxed">
              Une connaissance approfondie des environnements de production industriels et des contraintes de test
            </p>
          </div>
        </div>

        <div class="flex gap-4 scroll-reveal" style="--delay: 0.2s">
          <div class="flex-shrink-0 w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center text-2xl">
            ⚡
          </div>
          <div>
            <h3 class="text-slate-200 text-xl font-semibold mb-2">Performance optimale</h3>
            <p class="text-slate-400 text-sm leading-relaxed">
              Des applications légères et rapides conçues pour ne pas impacter vos systèmes de production
            </p>
          </div>
        </div>

        <div class="flex gap-4 scroll-reveal" style="--delay: 0.3s">
          <div class="flex-shrink-0 w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center text-2xl">
            🔒
          </div>
          <div>
            <h3 class="text-slate-200 text-xl font-semibold mb-2">Sécurité renforcée</h3>
            <p class="text-slate-400 text-sm leading-relaxed">
              Protocoles de sécurité stricts pour protéger vos données et vos processus critiques
            </p>
          </div>
        </div>

        <div class="flex gap-4 scroll-reveal" style="--delay: 0.4s">
          <div class="flex-shrink-0 w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center text-2xl">
            🎯
          </div>
          <div>
            <h3 class="text-slate-200 text-xl font-semibold mb-2">Sur mesure</h3>
            <p class="text-slate-400 text-sm leading-relaxed">
              Solutions personnalisées adaptées précisément à vos besoins et contraintes spécifiques
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Testimonials Section -->
    <div class="max-w-4xl w-full mx-auto mt-24 mb-20">
      <h2 class="text-4xl font-bold text-center text-slate-200 mb-16 scroll-reveal">Ils nous ont fait confiance</h2>
      <div class="space-y-6">
        <div class="bg-slate-800/50 backdrop-blur-sm border border-slate-700 rounded-xl p-8 scroll-reveal" style="--delay: 0.1s">
          <p class="text-slate-300 text-lg italic mb-4 leading-relaxed">
            "K-DEV m'a permis d'exposer mon entreprise en ligne avec un site vitrine élégant et performant. Le processus a été fluide et le résultat est au rendez-vous !"
          </p>
          <p class="text-blue-400 font-semibold">CEO - "ösleya b."</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { onMounted } from 'vue';

export default {
  name: 'HomePage',
  setup() {
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
      setupScrollAnimations();
    });

    return {};
  }
};
</script>

<style scoped>
@keyframes float {
  0%, 100% {
    transform: translateY(0) translateX(0);
    opacity: 0.3;
  }
  50% {
    transform: translateY(-20px) translateX(10px);
    opacity: 0.6;
  }
}

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

.animate-float {
  animation: float 3s ease-in-out infinite;
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
</style>