"use client";

import Image from "next/image";
import React, { useEffect, useRef, useState } from "react";
import Reveal from "../components/ui/Reveal";

/**
 * Home — Refonte A+ (Hero + Signal Board)
 * - No external deps
 * - No randomness (no hydration mismatch)
 * - Subtle parallax via CSS vars (mouse), disabled for prefers-reduced-motion
 */

const SIGNAL_PARTICLES = [
  { left: "10%", top: "22%", s: 10, o: 0.18, d: "18s", delay: "-3s" },
  { left: "14%", top: "66%", s: 14, o: 0.12, d: "22s", delay: "-9s" },
  { left: "22%", top: "38%", s: 8, o: 0.16, d: "16s", delay: "-6s" },
  { left: "32%", top: "18%", s: 9, o: 0.14, d: "20s", delay: "-11s" },
  { left: "36%", top: "72%", s: 12, o: 0.11, d: "24s", delay: "-7s" },
  { left: "48%", top: "28%", s: 13, o: 0.10, d: "26s", delay: "-12s" },
  { left: "56%", top: "62%", s: 9, o: 0.12, d: "19s", delay: "-5s" },
  { left: "64%", top: "20%", s: 8, o: 0.13, d: "21s", delay: "-10s" },
  { left: "72%", top: "46%", s: 11, o: 0.10, d: "27s", delay: "-14s" },
  { left: "82%", top: "30%", s: 12, o: 0.09, d: "25s", delay: "-8s" },
  { left: "86%", top: "70%", s: 14, o: 0.08, d: "30s", delay: "-15s" },
];

function PoSSSignalBoard() {
  return (
    <div className="relative overflow-hidden rounded-2xl border border-white/15 bg-white/[0.06] shadow-[0_26px_80px_rgba(0,0,0,0.28)] backdrop-blur-md">
      {/* top sheen */}
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(800px_circle_at_20%_10%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(700px_circle_at_85%_30%,rgba(255,255,255,0.10),transparent_60%)]" />

      {/* animated grid stream */}
      <div className="pointer-events-none absolute inset-0 opacity-[0.28] bg-[linear-gradient(to_right,rgba(255,255,255,0.22)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.22)_1px,transparent_1px)] bg-[size:42px_42px] noor-grid-stream" />

      {/* content */}
      <div className="relative p-6">
        <div className="flex items-start justify-between gap-4 mb-5">
          <div>
            <p className="text-xs uppercase tracking-wide text-white/70">
              PoSS Signal Board
            </p>
            <h3 className="text-lg font-semibold text-white mt-1">
              Verifiable Signals • Curator Validation
            </h3>
          </div>

          <div className="rounded-full border border-white/20 bg-white/10 px-3 py-1 text-xs text-white/80">
            Controlled
          </div>
        </div>

        <div className="relative rounded-xl border border-white/15 bg-black/10 p-4 overflow-hidden">
          {/* scan */}
          <div className="pointer-events-none absolute inset-0 noor-board-scan" />

          {/* animated lines + pulses */}
          <svg
            className="block w-full h-[190px]"
            viewBox="0 0 540 190"
            preserveAspectRatio="none"
            aria-hidden="true"
          >
            <defs>
              <linearGradient id="line" x1="0" y1="0" x2="1" y2="0">
                <stop offset="0" stopColor="rgba(255,255,255,0.06)" />
                <stop offset="0.5" stopColor="rgba(255,255,255,0.22)" />
                <stop offset="1" stopColor="rgba(255,255,255,0.06)" />
              </linearGradient>
              <linearGradient id="pulse" x1="0" y1="0" x2="1" y2="0">
                <stop offset="0" stopColor="rgba(255,255,255,0)" />
                <stop offset="0.45" stopColor="rgba(255,255,255,0.70)" />
                <stop offset="0.55" stopColor="rgba(255,255,255,0.70)" />
                <stop offset="1" stopColor="rgba(255,255,255,0)" />
              </linearGradient>
              <filter id="g" x="-30%" y="-30%" width="160%" height="160%">
                <feGaussianBlur stdDeviation="1.8" result="b" />
                <feMerge>
                  <feMergeNode in="b" />
                  <feMergeNode in="SourceGraphic" />
                </feMerge>
              </filter>
            </defs>

            {/* baseline layers */}
            <g stroke="url(#line)" strokeWidth="1" opacity="0.95" fill="none">
              <path d="M20 30 C120 10, 210 60, 310 42 C390 30, 450 18, 520 34" />
              <path d="M20 86 C120 66, 210 118, 310 96 C390 78, 450 70, 520 90" />
              <path d="M20 140 C120 126, 210 170, 310 150 C390 134, 450 120, 520 146" />
            </g>

            {/* flowing pulses */}
            <g
              filter="url(#g)"
              stroke="url(#pulse)"
              strokeWidth="2"
              opacity="0.70"
              fill="none"
            >
              <path
                className="noor-flow noor-flow-1"
                d="M20 30 C120 10, 210 60, 310 42 C390 30, 450 18, 520 34"
              />
              <path
                className="noor-flow noor-flow-2"
                d="M20 86 C120 66, 210 118, 310 96 C390 78, 450 70, 520 90"
              />
              <path
                className="noor-flow noor-flow-3"
                d="M20 140 C120 126, 210 170, 310 150 C390 134, 450 120, 520 146"
              />
            </g>

            {/* nodes */}
            <g filter="url(#g)">
              {[
                [88, 22],
                [170, 46],
                [250, 52],
                [330, 40],
                [420, 26],
                [492, 36],
                [92, 78],
                [178, 98],
                [260, 112],
                [338, 96],
                [430, 84],
                [505, 92],
                [86, 132],
                [176, 154],
                [260, 166],
                [340, 152],
                [430, 132],
                [506, 146],
              ].map(([cx, cy], i) => (
                <g key={i} className="noor-node">
                  <circle
                    cx={cx}
                    cy={cy}
                    r="2.6"
                    fill="rgba(255,255,255,0.78)"
                  />
                  <circle
                    cx={cx}
                    cy={cy}
                    r="10"
                    fill="rgba(255,255,255,0.06)"
                  />
                </g>
              ))}
            </g>
          </svg>

          {/* small legend */}
          <div className="mt-3 flex flex-wrap items-center gap-2 text-[11px] text-white/70">
            <span className="inline-flex items-center gap-2 rounded-full border border-white/15 bg-white/5 px-2 py-1">
              <span className="h-1.5 w-1.5 rounded-full bg-white/70" />
              Signal stream
            </span>
            <span className="inline-flex items-center gap-2 rounded-full border border-white/15 bg-white/5 px-2 py-1">
              <span className="h-1.5 w-1.5 rounded-full bg-white/50" />
              Curator checkpoints
            </span>
            <span className="inline-flex items-center gap-2 rounded-full border border-white/15 bg-white/5 px-2 py-1">
              <span className="h-1.5 w-1.5 rounded-full bg-white/40" />
              Auditability-first
            </span>
          </div>
        </div>

        {/* facts */}
        <div className="mt-5 grid grid-cols-2 gap-3 text-xs text-white/80">
          <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
            <div className="text-white/60">Network</div>
            <div className="font-medium text-white">Private mainnet-like</div>
          </div>
          <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
            <div className="text-white/60">Consensus</div>
            <div className="font-medium text-white">Permissioned BFT</div>
          </div>
          <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
            <div className="text-white/60">Layer</div>
            <div className="font-medium text-white">Sovereign EVM L1</div>
          </div>
          <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
            <div className="text-white/60">PoSS</div>
            <div className="font-medium text-white">Application-layer</div>
          </div>
        </div>
      </div>
    </div>
  );
}

function ProofOfLivenessPanel() {
  const [live, setLive] = useState<{
    chain_id?: string;
    leader_height?: number;
    observed_at?: string;
    uptime_seconds?: number;
  } | null>(null);

  useEffect(() => {
    let stopped = false;

    const load = async () => {
      try {
        const res = await fetch("/liveness.json", { cache: "no-store" });
        if (!res.ok) return;
        const j = await res.json();

        // Strict allowlist: keep only the four authorized fields.
        const next = {
          chain_id: typeof j?.chain_id === "string" ? j.chain_id : undefined,
          leader_height:
            typeof j?.leader_height === "number" ? j.leader_height : undefined,
          observed_at:
            typeof j?.observed_at === "string" ? j.observed_at : undefined,
          uptime_seconds:
            typeof j?.uptime_seconds === "number" ? j.uptime_seconds : undefined,
        };

        if (!stopped) setLive(next);
      } catch {
        // silent by design
      }
    };

    load();
    const t = window.setInterval(load, 30_000);
    return () => {
      stopped = true;
      window.clearInterval(t);
    };
  }, []);

  return (
    <div className="relative overflow-hidden rounded-2xl border border-white/15 bg-white/[0.06] shadow-[0_26px_80px_rgba(0,0,0,0.22)] backdrop-blur-md">
      {/* top sheen */}
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(900px_circle_at_20%_10%,rgba(255,255,255,0.16),transparent_55%),radial-gradient(900px_circle_at_85%_30%,rgba(255,255,255,0.08),transparent_60%)]" />

      {/* subtle grid */}
      <div className="pointer-events-none absolute inset-0 opacity-[0.24] bg-[linear-gradient(to_right,rgba(255,255,255,0.18)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.18)_1px,transparent_1px)] bg-[size:42px_42px] noor-grid-stream" />

      <div className="relative p-6">
        <div className="flex items-start justify-between gap-4 mb-5">
          <div>
            <p className="text-xs uppercase tracking-wide text-white/70">
              Proof-of-Liveness
            </p>
            <h3 className="text-lg font-semibold text-white mt-1">
              Minimal public signal (read-only)
            </h3>
          </div>

          <div className="rounded-full border border-white/20 bg-white/10 px-3 py-1 text-xs text-white/80">
            30s refresh
          </div>
        </div>

        <div className="relative rounded-xl border border-white/15 bg-black/10 p-4 overflow-hidden">
          <div className="pointer-events-none absolute inset-0 noor-board-scan" />

          <div className="relative grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3 text-xs text-white/80">
            <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
              <div className="text-white/60">chain_id</div>
              <div className="font-mono text-white truncate">
                {live?.chain_id ? live.chain_id : "—"}
              </div>
            </div>

            <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
              <div className="text-white/60">leader_height</div>
              <div className="font-mono text-white">
                {typeof live?.leader_height === "number" ? live.leader_height : "—"}
              </div>
            </div>

            <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
              <div className="text-white/60">observed_at</div>
              <div className="font-mono text-white truncate">
                {live?.observed_at ? live.observed_at : "—"}
              </div>
            </div>

            <div className="rounded-lg border border-white/12 bg-white/5 px-3 py-2">
              <div className="text-white/60">uptime_seconds</div>
              <div className="font-mono text-white">
                {typeof live?.uptime_seconds === "number" ? live.uptime_seconds : "—"}
              </div>
            </div>
          </div>

          <div className="relative mt-3 text-[11px] text-white/60">
            Public surface:{" "}
            <span className="font-mono text-white/70">/liveness.json</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default function HomePage() {
  const heroRef = useRef<HTMLElement | null>(null);

  useEffect(() => {
    const el = heroRef.current;
    if (!el) return;

    const media = window.matchMedia("(prefers-reduced-motion: reduce)");
    if (media.matches) return;

    const onMove = (e: MouseEvent) => {
      const r = el.getBoundingClientRect();
      const x = (e.clientX - r.left) / r.width; // 0..1
      const y = (e.clientY - r.top) / r.height; // 0..1
      // map to -1..1
      const mx = (x - 0.5) * 2;
      const my = (y - 0.5) * 2;
      el.style.setProperty("--mx", mx.toFixed(3));
      el.style.setProperty("--my", my.toFixed(3));
    };

    window.addEventListener("mousemove", onMove, { passive: true });
    return () => window.removeEventListener("mousemove", onMove);
  }, []);

  return (
    <main className="w-full">
      {/* HERO — stronger composition (text + signal board) */}
      <section
        ref={(n) => {
          heroRef.current = n;
        }}
        className="relative isolate w-full overflow-hidden text-white"
      >
        {/* Base gradient */}
        <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />

        {/* Spotlight + depth */}
        <div className="absolute inset-0 noor-spotlight" />

        {/* Subtle grid */}
        <div className="absolute inset-0 opacity-[0.12] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:56px_56px]" />

        {/* Particles (anchored, not “a single dot”) */}
        <div className="pointer-events-none absolute inset-0 noor-parallax-1">
          {SIGNAL_PARTICLES.map((p, i) => (
            <span
              key={i}
              className="noor-particle"
              style={
                {
                  left: p.left,
                  top: p.top,
                  width: `${p.s}px`,
                  height: `${p.s}px`,
                  opacity: p.o,
                  ["--dur" as any]: p.d,
                  ["--delay" as any]: p.delay,
                } as React.CSSProperties
              }
            />
          ))}
        </div>

        {/* Scan line */}
        <div className="pointer-events-none absolute inset-0 noor-scan" />

        <div className="relative">
          <div className="container py-14 sm:py-16 md:py-24">
            <div className="grid grid-cols-1 lg:grid-cols-12 gap-10 lg:gap-12 items-center">
              {/* LEFT — copy */}
              <div className="lg:col-span-7">
                <div className="flex items-center gap-3 mb-6">
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
                  Private mainnet-like environment — controlled operation, non-public by design.
                </p>

                <p className="text-sm sm:text-base md:text-lg text-white/92 leading-relaxed mb-8 max-w-2xl">
                  A Social Signal Blockchain powered by PoSS. Designed for transparent
                  participation, curator validation, and a fixed-supply digital model
                  free from financial speculation.
                </p>

                <div className="flex flex-col sm:flex-row flex-wrap gap-3 sm:gap-4">
                  <a
                    href="/technology"
                    className="w-full sm:w-auto text-center px-6 py-3 rounded-md text-sm md:text-base font-medium
                               bg-white text-[#0B1B3A] ring-1 ring-white/30
                               shadow-[0_14px_44px_rgba(0,0,0,0.22)]
                               transition-all duration-200 hover:-translate-y-0.5 hover:shadow-[0_18px_58px_rgba(0,0,0,0.28)]"
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

                {/* micro-proof line (serious, not marketing) */}
                <div className="mt-8 flex flex-wrap gap-2 text-xs text-white/75">
                  <span className="rounded-full border border-white/20 bg-white/10 px-3 py-1">
                    Tagged builds: M10 / M11 / M12.2
                  </span>
                  <span className="rounded-full border border-white/20 bg-white/10 px-3 py-1">
                    PoSS: application-layer (not consensus)
                  </span>
                  <span className="rounded-full border border-white/20 bg-white/10 px-3 py-1">
                    Legal-Light posture (no yield, no custody)
                  </span>
                </div>
              </div>

              {/* RIGHT — Signal Board */}
              <div className="lg:col-span-5 noor-parallax-2">
                <PoSSSignalBoard />
              </div>
            </div>
          </div>
        </div>

        {/* PROOF-OF-LIVENESS — included in hero background */}
        <div className="relative">
          <div className="container -mt-10 pb-10 md:-mt-14 md:pb-14">
            <ProofOfLivenessPanel />
          </div>
        </div>

        {/* Global CSS, scoped by noor-* classnames */}
        <style jsx global>{`
          @media (prefers-reduced-motion: reduce) {
            .noor-particle,
            .noor-scan,
            .noor-flow,
            .noor-node,
            .noor-grid-stream,
            .noor-board-scan,
            .noor-parallax-1,
            .noor-parallax-2,
            .noor-spotlight {
              animation: none !important;
              transform: none !important;
            }
          }

          /* spotlight + parallax */
          .noor-spotlight {
            background:
              radial-gradient(900px circle at 18% 18%, rgba(255,255,255,0.18), transparent 55%),
              radial-gradient(900px circle at 85% 25%, rgba(0,0,0,0.22), transparent 55%),
              radial-gradient(1200px circle at 50% 120%, rgba(0,0,0,0.32), transparent 62%);
            transform: translate3d(calc(var(--mx, 0) * 10px), calc(var(--my, 0) * 10px), 0);
            transition: transform 120ms linear;
            will-change: transform;
          }
          .noor-parallax-1 {
            transform: translate3d(calc(var(--mx, 0) * 8px), calc(var(--my, 0) * 6px), 0);
            transition: transform 120ms linear;
            will-change: transform;
          }
          .noor-parallax-2 {
            transform: translate3d(calc(var(--mx, 0) * -10px), calc(var(--my, 0) * -8px), 0);
            transition: transform 120ms linear;
            will-change: transform;
          }

          /* particles */
          .noor-particle {
            position: absolute;
            border-radius: 9999px;
            background: rgba(255, 255, 255, 0.92);
            box-shadow:
              0 0 0 1px rgba(255, 255, 255, 0.08),
              0 0 34px rgba(255, 255, 255, 0.14);
            animation: noor-drift var(--dur, 20s) ease-in-out infinite;
            animation-delay: var(--delay, 0s);
          }
          @keyframes noor-drift {
            0% { transform: translate3d(0,0,0) scale(1); }
            35% { transform: translate3d(12px,-12px,0) scale(1.05); }
            70% { transform: translate3d(-14px,10px,0) scale(0.98); }
            100% { transform: translate3d(0,0,0) scale(1); }
          }

          /* scan */
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
            animation: noor-scan 7.2s ease-in-out infinite;
          }
          @keyframes noor-scan {
            0% { transform: translateY(-55%); opacity: 0.25; }
            45% { opacity: 0.68; }
            100% { transform: translateY(85%); opacity: 0.25; }
          }

          /* board grid stream */
          .noor-grid-stream {
            background-position: 0 0;
            animation: noor-grid 10s linear infinite;
          }
          @keyframes noor-grid {
            0% { background-position: 0 0; }
            100% { background-position: 120px 80px; }
          }

          /* board scan */
          .noor-board-scan {
            background: linear-gradient(
              to right,
              transparent,
              rgba(255,255,255,0.10),
              transparent
            );
            opacity: 0.55;
            transform: translateX(-40%);
            animation: noor-board-scan 6.6s ease-in-out infinite;
          }
          @keyframes noor-board-scan {
            0% { transform: translateX(-55%); opacity: 0.22; }
            45% { opacity: 0.62; }
            100% { transform: translateX(85%); opacity: 0.22; }
          }

          /* flowing pulses on svg paths */
          .noor-flow {
            stroke-dasharray: 36 210;
            stroke-dashoffset: 0;
            animation: noor-flow 6.8s linear infinite;
          }
          .noor-flow-2 { animation-duration: 8.2s; animation-delay: -2.2s; }
          .noor-flow-3 { animation-duration: 9.6s; animation-delay: -3.6s; }
          @keyframes noor-flow {
            0% { stroke-dashoffset: 0; opacity: 0.45; }
            35% { opacity: 0.92; }
            100% { stroke-dashoffset: -520; opacity: 0.55; }
          }

          /* svg nodes pulse */
          .noor-node {
            transform-origin: center;
            animation: noor-node-pulse 3.6s ease-in-out infinite;
          }
          .noor-node:nth-child(2n) { animation-duration: 4.6s; }
          .noor-node:nth-child(3n) { animation-duration: 5.3s; }
          @keyframes noor-node-pulse {
            0% { opacity: 0.72; transform: scale(1); }
            50% { opacity: 1; transform: scale(1.08); }
            100% { opacity: 0.72; transform: scale(1); }
          }
        `}</style>
      </section>

      {/* STATUS + INTRO + PoSS FRAMING — layout validé (1 full + 2 half) */}
      <section className="container py-10 md:py-14">
        <div className="grid grid-cols-1 gap-6">
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
                  Application layer for governance and verifiable contribution signals (not consensus)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Public access:</span>{" "}
                  Limited by design until feature completeness and security review
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Proof-of-liveness:</span>{" "}
                  Public status endpoint <span className="font-mono">/liveness.json</span> (minimal signal only)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Phase 7 baseline:</span>{" "}
                  Frozen (controlled operation; no additional public surfaces)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Reference build:</span>{" "}
                  M10-MAINNETLIKE-STABLE / M11-DAPPS-STABLE / M12.2-WORLDSTATE-RPC-NONCE-BALANCE
                </li>
              </ul>
            </div>
          </Reveal>

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
                  infrastructure for curators, participants, institutions and communities.
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
