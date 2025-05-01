import Image from "next/image"
import Link from "next/link"

export default function Footer() {
  return (
    <footer className="bg-[#1e3a5f] text-white py-8">
      <div className="container mx-auto px-4">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-2">
          {/* Left side - Logo and description */}
          <div className="max-w-md">
            <div className="mb-4">
              <div className="bg-white rounded-lg p-2 inline-block">
                <div className="h-8 w-36 bg-white relative">
                  <Image
                    src="/LokerHub_Logo.svg"
                    alt="logo"
                    width={222}
                    height={42}
                    className="w-full h-auto"
                    priority
                  />
                </div>
              </div>
            </div>
            <p>
              Platform rekrutmen yang membantu pencari kerja menemukan peluang terbaik dan mempertemukan perekrut dalam
              menyaring kandidat yang sesuai. Dengan teknologi AI, kami mempercepat proses pencocokan CV dan menyediakan
              pengalaman rekrutmen yang lebih efisien.
            </p>
          </div>

          {/* Right side - Project info */}
          <div className="text-base">
            <h3 className="font-semibold mb-3 text-lg">Senior Project 2025</h3>
            <ul className="space-y-1">
              <li className="flex justify-between gap-2">
                <span>Muhammad Fachrizal Giffari</span>
                <span>22/504570/TK/55192</span>
              </li>
              <li className="flex justify-between gap-2">
                <span>Muhammad Rendy</span>
                <span>22/503275/TK/54980</span>
              </li>
              <li className="flex justify-between gap-2">
                <span>Nashatra Aqila Ramadhani</span>
                <span>22/504580/TK/55195</span>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </footer>
  )
}
