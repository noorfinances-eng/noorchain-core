import Image from "next/image";
import Reveal from "../components/ui/Reveal";

export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — premium depth (UI only, same text) */}
      <section className="relative isolate w-full overflow-hidden text-white">
        {/* Base gradient */}
        <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />

        {/* Depth: soft radial shading */}
        <div className="absolute inset-0 bg-[radial-gradient(900px_circle_at_18%_18%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(900px_circle_at_85%_25%,rgba(0,0,0,0.18),transparent_55%),radial-gradient(1100px_circle_at_50%_115%,rgba(0,0,0,0.28),transparent_62%)]" />

        {/* Subtle grid */}
        <div className="absolute inset-0 opacity-[0.16] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:52px_52px]" />

        {/* Animated orbs */}
        <div className="pointer-events-none absolute -top-24 -left-24 h-[420px] w-[420px] rounded-full bg-white/22 blur-3xl noor-anim noor-orb-1" />
        <div className="pointer-events-none absolute -bottom-28 left-1/3 h-[520px] w-[520px] rounded-full bg-white/16 blur-3xl noor-anim noor-orb-2" />
        <div className="pointer-events-none absolute top-10 -right-28 h-[460px] w-[460px] rounded-full bg-white/14 blur-3xl noor-anim noor-orb-3" />

        {/* Content */}
        <div className="relative">
          <div className="container py-14 sm:py-18 md:py-24">
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
      </section>

      {/* STATUS + INTRO + PoSS FRAMING — same content, reveal animation */}
      <section className="container py-10 md:py-14">
        <div className="max-w-3xl space-y-6">
          <Reveal delayMs={0}>
            <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm">
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

          <Reveal delayMs={80}>
            <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm">
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
            <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm">
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
      </section>
    </main>
  );
}
