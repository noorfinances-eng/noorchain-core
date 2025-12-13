export default function Footer() {
  return (
    <footer className="w-full bg-white border-t border-gray-soft mt-12 py-6 text-sm">
      <div className="container flex flex-col items-center text-center gap-3">

        {/* TITLE */}
        <div className="text-navy font-semibold text-base">
          NOORCHAIN — Social Signal Blockchain
        </div>

        {/* INTERNAL LINKS */}
        <div className="flex flex-wrap justify-center gap-4 text-gray-700">
          <a href="/legal" className="hover:text-primary transition">
            Legal Notices
          </a>
          <a href="/legal" className="hover:text-primary transition">
            Compliance & Risks
          </a>
          <a href="/legal" className="hover:text-primary transition">
            Privacy
          </a>
        </div>

        {/* EXTERNAL LINKS */}
        <div className="flex flex-wrap justify-center gap-4 text-gray-700">
          <a
            href="https://github.com/noorfinances-eng/noorchain-core"
            target="_blank"
            className="hover:text-primary transition"
          >
            GitHub
          </a>
          <a
            href="https://x.com/noorchainOrg"
            target="_blank"
            className="hover:text-primary transition"
          >
            X (Twitter)
          </a>
          <a
            href="mailto:contact@noorchain.io"
            className="hover:text-primary transition"
          >
            Contact
          </a>
        </div>

        {/* LINE */}
        <div className="h-px w-24 bg-gray-soft my-1" />

        {/* COPYRIGHT */}
        <div className="text-gray-600 text-xs">
          © {new Date().getFullYear()} NOORCHAIN. All rights reserved.
        </div>
      </div>
    </footer>
  );
}
