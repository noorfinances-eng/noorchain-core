export default function Footer() {
  return (
    <footer className="w-full bg-white border-t border-gray-soft mt-16 py-10 text-sm">
      <div className="container flex flex-col gap-6">

        {/* TITLE */}
        <div className="text-navy font-semibold text-base">
          NOORCHAIN — Social Signal Blockchain
        </div>

        {/* INTERNAL LINKS */}
        <div className="flex flex-wrap gap-5 text-gray-700">
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
        <div className="flex flex-wrap gap-5 text-gray-700">
          <a
            href="https://github.com"
            target="_blank"
            className="hover:text-primary transition"
          >
            GitHub
          </a>
          <a
            href="https://x.com"
            target="_blank"
            className="hover:text-primary transition"
          >
            X (Twitter)
          </a>
          <a
            href="mailto:contact@noorchain.org"
            className="hover:text-primary transition"
          >
            Contact
          </a>
        </div>

        {/* BOTTOM LINE */}
        <div className="h-px w-full bg-gray-soft" />

        {/* COPYRIGHT */}
        <div className="text-gray-600">
          © {new Date().getFullYear()} NOORCHAIN. All rights reserved.
        </div>
      </div>
    </footer>
  );
}
