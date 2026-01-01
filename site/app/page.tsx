import Image from "next/image";
import Reveal from "../components/ui/Reveal";
import type { CSSProperties } from "react";

/**
 * Home — Refonte A (Hero + PoSS signal mesh + signal board)
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
  { left: "90%", top: "36%", size: 14, o: 0.1, d: "25s", delay: "-13s" },
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

      <g opacity="0.9" stroke="url(#meshStroke)" strokeWidth="1">
        <path d="M80 120 L380 80 L640 160 L930 110 L1120 180" fill="none" />
        <path d="M120 280 L360 220 L640 260 L900 240 L1140 310" fill="none" />
        <path d="M200 420 L420 360 L690 400 L980 360 L1160 420" fill="none" />
        <path d="M260 60 L300 220 L360 420" fill="none" />
        <path d="M520 40 L520 210 L560 420" fill="none" />
        <path d="M860 50 L840 240 L820 430" fill="none" />
        <path d="M1060 90 L980 240 L900 420" fill="none" />
      </g>

      <g
        filter="url(#softGlow)"
        stroke="url(#pulseStroke)"
        strokeWidth="2"
        fill="none"
        opacity="0.65"
      >
        <path
          className="noor-pulse-line noor-pulse-line-1"
          d="M80 120 L380 80 L640 160 L930 110 L1120 180"
        />
        <path
          className="noor-pulse-line noor-pulse-line-2"
          d="M120 280 L360 220 L640 260 L900 240 L1140 310"
        />
        <path
          className="noor-pulse-line noor-pulse-line-3"
          d="M200 420 L420 360 L690 400 L980 360 L1160 420"
        />
      </g>

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
            <circle
              cx={n.cx}
              cy={n.cy}
              r={n.r * 3.2}
              fill="rgba(255,255,255,0.06)"
            />
          </g>
        ))}
      </g>
    </svg>
  );
}

function MiniSignalChart() {
  return (
    <svg
      className="absolute inset-0 h-full w-full opacity-[0.85]"
      viewBox="0 0 520 210"
      preserveAspectRatio="none"
      aria-hidden="true"
    >
      <defs>
        <linearGradient id="chartLine" x1="0" y1="0" x2="1" y2="0">
          <stop offset="0" stopColor="rgba(255,255,255,0.10)" />
          <stop offset="0.5" stopColor="rgba(255,255,255,0.40)" />
          <stop offset="1" stopColor="rgba(255,255,255,0.10)" />
        </linearGradient>
        <filter id="chartGlow" x="-40%" y="-40%" width="180%" height="180%">
          <feGaussianBlur stdDeviation="2" result="b" />
          <feMerge>
            <feMergeNode in="b" />
            <feMergeNode in="SourceGraphic" />
          </feMerge>
        </filter>
      </defs>

      {[
        "M18 52 C 90 26, 140 86, 210 56 S 330 72, 408 48 S 470 74, 502 60",
        "M18 104 C 92 84, 148 132, 214 108 S 332 126, 408 104 S 470 130, 502 116",
        "M18 156 C 96 138, 150 184, 218 162 S 334 178, 410 158 S 470 184, 502 170",
      ].map((d, i) => (
        <g key={i} filter="url(#chartGlow)">
          <path d={d} stroke="rgba(255,255,255,0.10)" strokeWidth="2" fill="none" />
          <path
            d={d}
            className={`noor-pulse-line ${i === 1 ? "noor-pulse-line-2" : i === 2 ? "noor-pulse-line-3" : "noor-pulse-line-1"}`}
            stroke="url(#chartLine)"
            strokeWidth="2"
            fill="none"
          />
        </g>
      ))}

      {[
        { x: 90, y: 40 },
        { x: 160, y: 78 },
        { x: 250, y: 66 },
        { x: 330, y: 98 },
        { x: 420, y: 58 },
        { x: 470, y: 92 },
      ].map((p, i) => (
        <g key={i} filter="url(#chartGlow)">
          <circle cx={p.x} cy={p.y + 10} r="4" fill="rgba(255,255,255,0.65)" />
          <circle cx={p.x} cy={p.y + 10} r="12" fill="rgba(255,255,255,0.06)" />
        </g>
      ))}
    </svg>
  );
}

function PoSSSignalBoard() {
  return (
    <div
      className="
        relative overflow-hidden rounded-2xl border border-white/15 bg-white/[0.06]
        shadow-[0_26px_80px_rgba(0,0,0,0.28)] backdrop-blur-md
        w-[560px] h-[300px]
      "
      aria-hidden="true"
    >
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(800px_circle_at_20%_18%,rgba(255,255,255,0.20),transparent_60%)]" />
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(900px_circle_at_88%_70%,rgba(0,0,0,0.26),transparent_60%)]" />

      <div className="pointer-events-none absolute inset-0">
        <span className="noor-beam noor-beam-1" />
        <span className="noor-beam noor-beam-2" />
        <span className="noor-beam noor-beam-3" />
        <span className="noor-beam noor-beam-4" />
      </div>

      {/* Header */}
      <div className="relative px-7 pt-6">
        <div className="flex items-center justify-between">
          <p className="text-[11px] tracking-[0.22em] uppercase text-white/70">
            PoSS Signal Board
          </p>
          <span className="inline-flex items-center rounded-full border border-white/20 bg-white/10 px-4 py-1 text-xs text-white/80">
            Controlled
          </span>
        </div>

        <h3 className="mt-2 text-[26px] leading-[1.12] font-semibold text-white">
          Verifiable Signals •<br />
          Curator Validation
        </h3>
      </div>

      {/* Middle panel (compressed to fit 300px total height) */}
      <div className="relative mt-4 px-7">
        <div className="relative h-[150px] rounded-2xl bg-black/10 ring-1 ring-white/10 overflow-hidden">
          <MiniSignalChart />
        </div>

        <div className="mt-3 flex flex-wrap gap-2">
          <span className="inline-flex items-center gap-2 rounded-full border border-white/18 bg-white/10 px-4 py-2 text-xs text-white/80">
            <span className="h-2 w-2 rounded-full bg-white/70" />
            Signal stream
          </span>
          <span className="inline-flex items-center gap-2 rounded-full border border-white/18 bg-white/10 px-4 py-2 text-xs text-white/80">
            <span className="h-2 w-2 rounded-full bg-white/70" />
            Curator checkpoints
          </span>
          <span className="inline-flex items-center gap-2 rounded-full border border-white/18 bg-white/10 px-4 py-2 text-xs text-white/80">
            <span className="h-2 w-2 rounded-full bg-white/70" />
            Auditability-first
          </span>
        </div>
      </div>

      {/* Bottom grid (kept, but tighter) */}
      <div className="absolute bottom-4 left-0 right-0 px-7">
        <div className="grid grid-cols-2 gap-3">
          {[
            ["Network", "Private mainnet-like"],
            ["Consensus", "Permissioned BFT"],
            ["Layer", "Sovereign EVM L1"],
            ["PoSS", "Application-layer"],
          ].map(([k, v]) => (
            <div
              key={k}
              className="rounded-xl border border-white/16 bg-white/[0.06] px-4 py-3"
            >
              <p className="text-xs text-white/55">{k}</p>
              <p className="text-sm font-semibold text-white/90">{v}</p>
            </div>
          ))}
        </div>
      </div>

      <div className="pointer-events-none absolute inset-0 ring-1 ring-white/10" />
    </div>
  );
}

export default function HomePage() {
  return (
    <main className="w-full">
      <section className="relative isolate w-full overflow-hidden text-white">
        <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />
        <div className="absolute inset-0 bg-[radial-gradient(900px_circle_at_18%_18%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(900px_circle_at_85%_25%,rgba(0,0,0,0.20),transparent_55%),radial-gradient(1100px_circle_at_50%_115%,rgba(0,0,0,0.30),transparent_62%)]" />
        <div className="absolute inset-0 opacity-[0.14] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:52px_52px]" />

        <div className="absolute inset-0 opacity-[0.85]">
          <SignalMesh />
        </div>

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
                } as CSSProperties
              }
            />
          ))}
        </div>

        {/* Board placement: slightly more right + lower, and never higher than headline */}
        <div className="pointer-events-none absolute right-[-90px] top-[96px] hidden lg:block">
          <PoSSSignalBoard />
        </div>

        <div className="pointer-events-none absolute inset-0 noor-scan" />

        <div className="relative">
          <div className="container py-16 sm:py-18 md:py-24">
            <div className="grid grid-cols-1 lg:grid-cols-[minmax(0,1fr)_560px] gap-10 items-start">
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
                    className="w-full sm:w-auto text-center h-11 inline-flex items-center justify-center px-6 rounded-md text-sm md:text-base font-medium
                               bg-white text-[#0B1B3A] ring-1 ring-white/30
                               shadow-[0_14px_44px_rgba(0,0,0,0.22)]
                               transition-all duration-200 hover:-translate-y-0.5 hover:shadow-[0_18px_58px_rgba(0,0,0,0.28)]"
                  >
                    Explore Technology
                  </a>

                  <a
                    href="/genesis"
                    className="w-full sm:w-auto text-center h-11 inline-flex items-center justify-center px-6 rounded-md text-sm md:text-base font-medium
                               border border-white/70 text-white
                               bg-white/0 transition-all duration-200
                               hover:bg-white/10 hover:-translate-y-0.5"
                  >
                    Genesis Overview
                  </a>
                </div>
              </div>

              <div className="hidden lg:block" />
            </div>
          </div>
        </div>
      </section>

      <section className="container py-10 md:py-14">
        <div className="grid grid-cols-1 gap-6">
          <Reveal delayMs={0}>
            <div className="rounded-xl border border-gray-200 bg-white p-7 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                <h2 className="text-xl font-bold text-gray-900">Current Project Status</h2>
                <p className="text-xs sm:text-sm text-gray-500">Last updated: 2026-01-01</p>
              </div>

              <ul className="text-sm text-gray-700 leading-relaxed space-y-2">
                <li><span className="font-semibold text-gray-900">Network:</span> Private mainnet-like environment (continuous operation)</li>
                <li><span className="font-semibold text-gray-900">Consensus:</span> Permissioned BFT</li>
                <li><span className="font-semibold text-gray-900">Layer 1:</span> Sovereign EVM L1</li>
                <li><span className="font-semibold text-gray-900">PoSS:</span> Application layer for governance and verifiable contribution signals (not consensus)</li>
                <li><span className="font-semibold text-gray-900">Public access:</span> Limited by design until feature completeness and security review</li>
                <li><span className="font-semibold text-gray-900">Reference build:</span> M10-MAINNETLIKE-STABLE / M11-DAPPS-STABLE / M12.2-WORLDSTATE-RPC-NONCE-BALANCE</li>
              </ul>
            </div>
          </Reveal>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <Reveal delayMs={80}>
              <div className="rounded-xl border border-gray-200 bg-white p-7 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 md:min-h-[320px]">
                <h2 className="text-xl font-bold text-gray-900 mb-4">A New Approach to Blockchain Design</h2>
                <p className="text-sm text-gray-700 leading-relaxed">
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
              <div className="rounded-xl border border-gray-200 bg-white p-7 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 md:min-h-[320px]">
                <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                  <h2 className="text-xl font-bold text-gray-900">PoSS Framing (Non-Consensus)</h2>
                  <a href="/poss" className="text-sm font-medium text-blue-700 hover:text-blue-900 transition">
                    Read PoSS framing
                  </a>
                </div>

                <ul className="text-sm text-gray-700 leading-relaxed space-y-2">
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
