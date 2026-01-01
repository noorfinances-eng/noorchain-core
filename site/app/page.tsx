import Image from "next/image";
import Reveal from "../components/ui/Reveal";

/**
 * Home — Refonte A (Hero + signal mesh)
 * - No external deps
 * - No randomness (no hydration mismatch)
 * - Animations are CSS-based + respects prefers-reduced-motion
 */

const SIGNAL_PARTICLES = [
  { left: "6%", top: "18%", size: 10, o: 0.22, d: "14s", delay: "0s" },
  { left: "10%", top: "62%", size: 14, o: 0.16, d: "18s", delay: "-4s" },
  { left: "18%", top: "34%", size: 8, o: 0.18, d: "12s", delay: "-2s" },
  { left: "22%", top: "78%", size: 12, o: 0.14, d: "20s", delay: "-9s" },
  { left: "28%", top: "22%", size: 9, o: 0.17, d: "16s", delay: "-6s" },
  { left: "34%", top: "52%", size: 16, o: 0.12, d: "22s", delay: "-10s" },
  { left: "40%", top: "14%", size: 7, o: 0.18, d: "15s", delay: "-3s" },
  { left: "44%", top: "70%", size: 11, o: 0.13, d: "19s", delay: "-8s" },
  { left: "52%", top: "30%", size: 13, o: 0.13, d: "21s", delay: "-11s" },
  { left: "58%", top: "58%", size: 9, o: 0.15, d: "17s", delay: "-5s" },
  { left: "64%", top: "20%", size: 8, o: 0.16, d: "14s", delay: "-7s" },
  { left: "68%", top: "74%", size: 15, o: 0.12, d: "24s", delay: "-12s" },
  { left: "74%", top: "40%", size: 10, o: 0.14, d: "18s", delay: "-6s" },
  { left: "80%", top: "16%", size: 12, o: 0.12, d: "20s", delay: "-9s" },
  { left: "84%", top: "66%", size: 8, o: 0.16, d: "16s", delay: "-2s" },
  { left: "90%", top: "36%", size: 14, o: 0.10, d: "25s", delay: "-13s" },
];

function SignalMesh() {
  return (
    <svg
      className="pointer-events-none absolute inset-0 h-full w-full"
      viewBox="0 0 1200 520"
      preserveAspectRatio="none"
      aria-hidden="true"
    >
      <defs>
        <linearGradient id="meshStroke" x1="0" y1="0" x2="1" y2="1">
          <stop offset="0" stopColor="rgba(255,255,255,0.16)" />
          <stop offset="1" stopColor="rgba(255,255,255,0.04)" />
        </linearGradient>

        <linearGradient id="pulseStroke" x1="0" y1="0" x2="1" y2="0">
          <stop offset="0" stopColor="rgba(255,255,255,0)" />
          <stop offset="0.45" stopColor="rgba(255,255,255,0.55)" />
          <stop offset="0.55" stopColor="rgba(255,255,255,0.55)" />
          <stop offset="1" stopColor="rgba(255,255,255,0)" />
        </linearGradient>

        <filter id="softGlow" x="-40%" y="-40%" width="180%" height="180%">
          <feGaussianBlur stdDeviation="2.2" result="b" />
          <feColorMatrix
            in="b"
            type="matrix"
            values="
              1 0 0 0 0
              0 1 0 0 0
              0 0 1 0 0
              0 0 0 0.8 0"
          />
          <feMerge>
            <feMergeNode />
            <feMergeNode in="SourceGraphic" />
          </feMerge>
        </filter>
      </defs>

      {/* Static mesh lines */}
      <g opacity="0.9" stroke="url(#meshStroke)" strokeWidth="1">
        <path d="M80 120 L380 80 L640 160 L930 110 L1120 180" fill="none" />
        <path d="M120 280 L360 220 L640 260 L900 240 L1140 310" fill="none" />
        <path d="M200 420 L420 360 L690 400 L980 360 L1160 420" fill="none" />
        <path d="M260 60 L300 220 L360 420" fill="none" />
        <path d="M520 40 L520 210 L560 420" fill="none" />
        <path d="M860 50 L840 240 L820 430" fill="none" />
        <path d="M1060 90 L980 240 L900 420" fill="none" />
      </g>

      {/* Pulses flowing over some lines */}
      <g filter="url(#softGlow)" stroke="url(#pulseStroke)" strokeWidth="2" fill="none" opacity="0.65">
        <path className="noor-pulse-line noor-pulse-line-1" d="M80 120 L380 80 L640 160 L930 110 L1120 180" />
        <path className="noor-pulse-line noor-pulse-line-2" d="M120 280 L360 220 L640 260 L900 240 L1140 310" />
        <path className="noor-pulse-line noor-pulse-line-3" d="M200 420 L420 360 L690 400 L980 360 L1160 420" />
      </g>

      {/* Nodes */}
      <g filter="url(#softGlow)">
        {[
          { cx: 80, cy: 120, r: 4 },
          { cx: 380, cy: 80, r: 4 },
          { cx: 640, cy: 160, r: 4 },
          { cx: 930, cy: 110, r: 4 },
          { cx: 1120, cy: 180, r: 4 },
          { cx: 120, cy: 280, r: 4 },
          { cx: 360, cy: 220, r: 4 },
          { cx: 640, cy: 260, r: 4 },
          { cx: 900, cy: 240, r: 4 },
          { cx: 1140, cy: 310, r: 4 },
          { cx: 200, cy: 420, r: 4 },
          { cx: 420, cy: 360, r: 4 },
          { cx: 690, cy: 400, r: 4 },
          { cx: 980, cy: 360, r: 4 },
          { cx: 1160, cy: 420, r: 4 },
          { cx: 300, cy: 220, r: 3 },
          { cx: 520, cy: 210, r: 3 },
          { cx: 840, cy: 240, r: 3 },
        ].map((n, i) => (
          <g key={i} className="noor-node">
            <circle cx={n.cx} cy={n.cy} r={n.r} fill="rgba(255,255,255,0.65)" />
            <circle cx={n.cx} cy={n.cy} r={n.r * 3.2} fill="rgba(255,255,255,0.06)" />
          </g>
        ))}
      </g>
    </svg>
  );
}

export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — strong visual identity (serious, PoSS-inspired) */}
      <section className="relative isolate w-full overflow-hidden text-white">
        {/* Base gradient */}
        <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />

        {/* Depth shading */}
        <div className="absolute inset-0 bg-[radial-gradient(900px_circle_at_18%_18%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(900px_circle_at_85%_25%,rgba(0,0,0,0.20),transparent_55%),radial-gradient(1100px_circle_at_50%_115%,rgba(0,0,0,0.30),transparent_62%)]" />

        {/* Subtle grid */}
        <div className="absolute inset-0 opacity-[0.14] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:52px_52px]" />

        {/* PoSS Signal Mesh */}
        <div className="absolute inset-0 opacity-[0.85]">
          <SignalMesh />
        </div>

        {/* Particles */}
        <div className="pointer-events-none absolute inset-0">
          {SIGNAL_PARTICLES.map((p, i) => (
            <span
              key={i}
              className="noor-particle"
              style={
                {
                  left: p.left,
                  top: p.top,
                  width: `${p.size}px`,
                  height: `${p.size}px`,
                  opacity: p.o,
                  ["--dur" as any]: p.d,
                  ["--delay" as any]: p.delay,
                } as React.CSSProperties
              }
            />
          ))}
        </div>

        {/* NOOR halo / rings (abstract, not a globe) */}
        <div className="pointer-events-none absolute right-[-140px] top-1/2 -translate-y-1/2">
          <div className="noor-ring noor-ring-1" />
          <div className="noor-ring noor-ring-2" />
          <div className="noor-ring noor-ring-3" />
        </div>

        {/* Scan line */}
        <div className="pointer-events-none absolute inset-0 noor-scan" />

        {/* Content */}
        <div className="relative">
          <div className="container py-14 sm:py-16 md:py-24">
            <div className="max-w-3xl">
              <div className="flex flex-col sm:flex-row items-start sm:items-center gap-3 sm:gap-4 mb-5">
                <div className="relative">
                  <div className="absolute -inset-3 rounded-2xl bg-white/15 blur-xl" />
                  <div className="relative rounded-2xl bg-white/10 ring-1 ring-white/25 p-2">
                    <Image
                      src="/logo-white.svg"
                      alt="NOORCHAIN Logo"
                      width={52}
                      height={52}
                      priority
                    />
                  </div>
                </div>

                <h1 className="text-3xl sm:text-4xl md:text-6xl font-extrabold tracking-tight">
                  NOORCHAIN
                </h1>
              </div>

              <p className="text-base sm:text-lg md:text-xl text-white/90 mb-3 font-medium">
                A Human-Centered Blockchain for Social Signals
              </p>

              <p className="text-xs sm:text-sm text-white/80 mb-6 font-medium">
                Private mainnet-like environment — controlled operation, non-public
                by design.
              </p>

              <p className="text-sm sm:text-base md:text-lg text-white/90 leading-relaxed mb-8">
                A Social Signal Blockchain powered by PoSS. Designed for transparent
                participation, curator validation, and a fixed-supply digital model
                free from financial speculation.
              </p>

              <div className="flex flex-col sm:flex-row flex-wrap gap-3 sm:gap-4">
                <a
                  href="/technology"
                  className="w-full sm:w-auto text-center px-6 py-3 rounded-md text-sm md:text-base font-medium
                             bg-white text-[#0B1B3A] ring-1 ring-white/30
                             shadow-[0_12px_34px_rgba(0,0,0,0.18)]
                             transition-all duration-200 hover:-translate-y-0.5 hover:shadow-[0_16px_46px_rgba(0,0,0,0.22)]"
                >
                  Explore Technology
                </a>

                <a
                  href="/genesis"
                  className="w-full sm:w-auto text-center px-6 py-3 rounded-md text-sm md:text-base font-medium
                             border border-white/70 text-white
                             bg-white/0 transition-all duration-200
                             hover:bg-white/10 hover:-translate-y-0.5"
                >
                  Genesis Overview
                </a>
              </div>
            </div>
          </div>
        </div>

        {/* Inline global CSS — scoped by classnames, doesn't touch the rest of the site */}
        <style jsx global>{`
          /* ---------- Motion safety ---------- */
          @media (prefers-reduced-motion: reduce) {
            .noor-particle,
            .noor-ring,
            .noor-scan,
            .noor-node,
            .noor-pulse-line {
              animation: none !important;
              transform: none !important;
            }
          }

          /* ---------- Particles ---------- */
          .noor-particle {
            position: absolute;
            border-radius: 9999px;
            background: rgba(255, 255, 255, 0.9);
            filter: blur(0.2px);
            box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.08),
              0 0 28px rgba(255, 255, 255, 0.14);
            animation: noor-drift var(--dur, 18s) ease-in-out infinite;
            animation-delay: var(--delay, 0s);
            will-change: transform;
          }
          @keyframes noor-drift {
            0% {
              transform: translate3d(0, 0, 0) scale(1);
            }
            35% {
              transform: translate3d(10px, -12px, 0) scale(1.06);
            }
            70% {
              transform: translate3d(-12px, 8px, 0) scale(0.98);
            }
            100% {
              transform: translate3d(0, 0, 0) scale(1);
            }
          }

          /* ---------- Mesh nodes ---------- */
          .noor-node {
            transform-origin: center;
            animation: noor-node-pulse 3.2s ease-in-out infinite;
          }
          .noor-node:nth-child(2n) {
            animation-duration: 4.4s;
          }
          .noor-node:nth-child(3n) {
            animation-duration: 5.1s;
          }
          @keyframes noor-node-pulse {
            0% {
              opacity: 0.75;
              transform: scale(1);
            }
            50% {
              opacity: 1;
              transform: scale(1.08);
            }
            100% {
              opacity: 0.75;
              transform: scale(1);
            }
          }

          /* ---------- Pulses flowing on lines ---------- */
          .noor-pulse-line {
            stroke-dasharray: 36 220;
            stroke-dashoffset: 0;
            animation: noor-flow 6.8s linear infinite;
          }
          .noor-pulse-line-2 {
            animation-duration: 8.2s;
            animation-delay: -2.4s;
          }
          .noor-pulse-line-3 {
            animation-duration: 9.6s;
            animation-delay: -4.2s;
          }
          @keyframes noor-flow {
            0% {
              stroke-dashoffset: 0;
              opacity: 0.45;
            }
            35% {
              opacity: 0.9;
            }
            100% {
              stroke-dashoffset: -520;
              opacity: 0.55;
            }
          }

          /* ---------- Rings / halo ---------- */
          .noor-ring {
            position: absolute;
            border-radius: 9999px;
            border: 1px solid rgba(255, 255, 255, 0.16);
            box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.05) inset,
              0 0 120px rgba(0, 0, 0, 0.22);
            background: radial-gradient(
              closest-side,
              rgba(255, 255, 255, 0.08),
              rgba(255, 255, 255, 0.03),
              transparent 70%
            );
            filter: blur(0.2px);
            animation: noor-rotate 22s linear infinite;
            will-change: transform;
          }
          .noor-ring-1 {
            width: 560px;
            height: 560px;
            opacity: 0.55;
          }
          .noor-ring-2 {
            width: 420px;
            height: 420px;
            opacity: 0.45;
            animation-duration: 28s;
            animation-direction: reverse;
            margin: 70px 0 0 70px;
          }
          .noor-ring-3 {
            width: 300px;
            height: 300px;
            opacity: 0.35;
            animation-duration: 34s;
            margin: 130px 0 0 130px;
          }
          @keyframes noor-rotate {
            0% {
              transform: rotate(0deg);
            }
            100% {
              transform: rotate(360deg);
            }
          }

          /* ---------- Scan line ---------- */
          .noor-scan {
            background: linear-gradient(
              to bottom,
              transparent,
              rgba(255, 255, 255, 0.06),
              transparent
            );
            mix-blend-mode: overlay;
            opacity: 0.55;
            transform: translateY(-40%);
            animation: noor-scan 7.5s ease-in-out infinite;
          }
          @keyframes noor-scan {
            0% {
              transform: translateY(-55%);
              opacity: 0.25;
            }
            45% {
              opacity: 0.65;
            }
            100% {
              transform: translateY(85%);
              opacity: 0.25;
            }
          }
        `}</style>
      </section>

      {/* STATUS + INTRO + PoSS FRAMING — 1 full-width + 2 half-width (layout validé) */}
      <section className="container py-10 md:py-14">
        <div className="grid grid-cols-1 gap-6">
          {/* ROW 1: FULL WIDTH */}
          <Reveal delayMs={0}>
            <div
              className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                         transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
            >
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                  Current Project Status
                </h2>
                <p className="text-xs sm:text-sm text-gray-500">
                  Last updated: 2026-01-01
                </p>
              </div>

              <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
                <li>
                  <span className="font-semibold text-gray-900">Network:</span>{" "}
                  Private mainnet-like environment (continuous operation)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Consensus:</span>{" "}
                  Permissioned BFT
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Layer 1:</span>{" "}
                  Sovereign EVM L1
                </li>
                <li>
                  <span className="font-semibold text-gray-900">PoSS:</span>{" "}
                  Application layer for governance and verifiable contribution
                  signals (not consensus)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Public access:</span>{" "}
                  Limited by design until feature completeness and security review
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Reference build:</span>{" "}
                  M10-MAINNETLIKE-STABLE / M11-DAPPS-STABLE / M12.2-WORLDSTATE-RPC-NONCE-BALANCE
                </li>
              </ul>
            </div>
          </Reveal>

          {/* ROW 2: TWO COLUMNS (md+) */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <Reveal delayMs={80}>
              <div
                className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                           transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
              >
                <h2 className="text-lg sm:text-xl font-bold text-gray-900 mb-4">
                  A New Approach to Blockchain Design
                </h2>

                <p className="text-sm sm:text-base text-gray-700 leading-relaxed">
                  NOORCHAIN introduces a mission-driven blockchain architecture
                  focused on verified social participation rather than financial
                  speculation. Powered by the PoSS protocol and aligned with Legal
                  Light CH, it provides a transparent and sustainable digital
                  infrastructure for curators, participants, institutions and
                  communities.
                </p>
              </div>
            </Reveal>

            <Reveal delayMs={140}>
              <div
                className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                           transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
              >
                <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                  <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                    PoSS Framing (Non-Consensus)
                  </h2>
                  <a
                    href="/poss"
                    className="text-sm font-medium text-blue-700 hover:text-blue-900 transition"
                  >
                    Read PoSS framing
                  </a>
                </div>

                <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
                  <li>
                    <span className="font-semibold text-gray-900">PoSS is not consensus.</span>{" "}
                    Network security is provided by a permissioned BFT consensus layer.
                  </li>
                  <li>
                    <span className="font-semibold text-gray-900">PoSS is an application layer.</span>{" "}
                    It structures governance, coordination, and verifiable contribution signals.
                  </li>
                  <li>
                    <span className="font-semibold text-gray-900">Economic posture:</span>{" "}
                    NUR is the native gas token; NOORCHAIN does not offer returns or custody.
                  </li>
                </ul>
              </div>
            </Reveal>
          </div>
        </div>
      </section>
    </main>
  );
}
