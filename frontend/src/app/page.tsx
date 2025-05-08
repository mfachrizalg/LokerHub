"use client"

import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { PiCheckSquare } from "react-icons/pi";
import { useRouter } from "next/navigation"
import Footer from '@/components/global/footer';
import Link from 'next/link';


export default function Home() {
  const router = useRouter()
  return (
    <div className="min-h-screen flex flex-col bg-white">
      <header className="border-b">
        <div className="container mx-auto py-6 flex items-center justify-between">
          <div className="flex items-center">
          <Image
                src="/LokerHub_Logo.svg"
                alt="logo"
                width={222}
                height={42}
                className="w-full "
              />
          </div>
          {/* Navigation Links */}
          <nav className="flex items-center gap-8">
            <div className="text-[#1e3a5f] font-medium">
              Kandidat
            </div>
            <div className="text-[#1e3a5f] font-medium">
              Rekruter
            </div>
            <Link
              href="/auth/login"
              className="bg-[#1e3a5f] text-white font-medium py-2 px-4 rounded hover:opacity-90"
            >
              Masuk
            </Link>
          </nav>
        </div>
      </header>

      <main className="flex-1">
        <section className="py-12 md:py-16 bg-[#f9f5f0]">
          <div className="container mx-auto">
            <div className="grid md:grid-cols-2 items-center">
              <div>
                <h1 className="text-4xl md-text-3xl font-bold text-blue-900 mb-4">Temukan Pekerjaan & Kandidat Terbaik!</h1>
                <p className="text-2xl text-black mb-8">Kami menghadirkan solusi berbasis AI yang membantu kandidat dalam menemukan pekerjaan yang paling sesuai dengan keterampilan dan pengalaman mereka, sekaligus mempermudah perusahaan dalam menemukan talenta terbaik dengan lebih cepat dan akurat.</p>
                <div className="flex flex-col sm:flex-row gap-8">
                  <Button 
                    className="w-50 h-10 bg-[#1e3a5f] hover:bg-[#1e3a5f]/90 text-white text-lg cursor-pointer"
                    onClick={() => router.push("/auth/register")}> Registrasi </Button>
                  <Button 
                    className="w-50 h-10 bg-white border-[#1e3a5f] text-[#1e3a5f] hover:bg-gray-100 text-lg cursor-pointer"
                    onClick={() => router.push("/auth/login")}>Masuk</Button>
                </div>
              </div>
              <div className="flex justify-end">
                <div className="relative rounded-full overflow-hidden">
                  <Image 
                    src="/HandsOnKeyboard.svg"
                    alt="Person typing on laptop"
                    width={325}
                    height={325}
                    className="object-cover">
                  </Image>
                </div>
              </div>
            </div>
          </div>

        </section>

        <section className="py-16 bg-white">
          <div className="container mx-auto">
            <h2 className="text-3xl md:text-4xl font-bold text-center text-[#1e3a5f] mb-12">
              Bergabung Sekarang & Temukan Kesempatan Baru!
            </h2>

            <div className="grid md:grid-cols-[1fr_2fr] gap-20 max-w-7xl mx-auto">
              <div className="flex flex-col justify-center">
                <h3 className="text-xl md:text-2xl font-bold text-[#1e3a5f]">KENAPA HARUS</h3>
                <Image 
                  src="/LokerHub_Logo.svg"
                  alt="logo"
                  width={132}
                  height={25}
                  className="w-auto max-w-full h-auto" />
              </div>

              <div className="space-y-7">
                <div className="border rounded-xl p-6 flex items-start gap-4 hover:shadow-md transition-shadow">
                  <div className="p-2 rounded-lg">
                    <PiCheckSquare className="h-8 w-8 text-[#1e3a5f]" />
                  </div>
                  <div>
                    <h3 className="font-bold text-[#1e3a5f] text-xl mb-1">
                      <span>Lebih Cepat, Lebih Tepat!</span>
                      <span className="font-normal"> – </span>
                      <span className="font-normal">
                        Temukan pekerjaan atau kandidat yang sesuai dalam waktu lebih singkat.
                      </span>
                    </h3>
                  </div>
                </div>

                <div className="border rounded-xl p-6 flex items-start gap-4 hover:shadow-md transition-shadow">
                  <div className="p-2 rounded-lg">
                    <PiCheckSquare className="h-8 w-8 text-[#1e3a5f]" />
                  </div>
                  <div>
                    <h3 className="font-bold text-[#1e3a5f] text-xl mb-1">
                      <span>Proses Seleksi Lebih Terarah</span>
                      <span className="font-normal"> – </span>
                      <span className="font-normal">
                        Dibantu dengan teknologi AI untuk mencocokkan kandidat dengan lowongan.
                      </span>
                    </h3>
                  </div>
                </div>

                <div className="border rounded-xl p-6 flex items-start gap-4 hover:shadow-md transition-shadow">
                  <div className="p-2 rounded-lg">
                    <PiCheckSquare className="h-8 w-8 text-[#1e3a5f]" />
                  </div>
                  <div>
                    <h3 className="font-bold text-[#1e3a5f] text-xl mb-1">
                      <span>Tanpa Ribet!</span>
                      <span className="font-normal"> – </span>
                      <span className="font-normal">
                        Cukup unggah CV atau posting lowongan, dan biarkan sistem bekerja untuk Anda.
                      </span>
                    </h3>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section className="py-16">
          <div className="container mx-auto">
            <div className="absolute border-2 p-4 rounded-lg mx-auto border-black relative mb-10">
              <div className="space-y-8 max-w-xs mx-auto">
                {/* Keuntungan Bagi Pelamar */}
                <div className="absolute -top-5 left-1/2 transform -translate-x-1/2 bg-gray-100 border-black px-6 rounded-xl border">
                  <h3 className="text-xl font-bold text-center text-[#1e3a5f] py-2 rounded-lg">
                    Keuntungan Bagi Pelamar
                  </h3>
                </div>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-3 gap-8 my-8 px-20">
                {/* Card 1 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md inline-block">
                      01
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Lamar Pekerjaan dengan Praktis
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Cukup unggah CV satu kali dan gunakan untuk melamar berbagai lowongan.
                  </p>
                </div>

                {/* Card 2 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md inline-block">
                      02
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Pantau Status Lamaran
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Tidak perlu menebak apakah lamaran Anda sudah diproses, cek statusnya langsung di LokerHub.
                  </p>
                </div>

                {/* Card 3 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md inline-block">
                      03
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Hemat Waktu & Lebih Efisien
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Fokus pada lowongan yang sesuai dengan keahlian dan pengalaman Anda.
                  </p>
                </div>
              </div>
            </div>
            
            <div className="border-2 py-4 rounded-lg mx-auto border-black relative mt-10">
              <div className="space-y-8 max-w-xs mx-auto">
                {/* Keuntungan Bagi Rekruter */}
                <div className="absolute -top-5 left-1/2 transform -translate-x-1/2 bg-gray-100 border-black px-6 rounded-xl border">
                  <h3 className="text-xl font-bold text-center text-[#1e3a5f] py-2 rounded-lg">
                    Keuntungan Bagi Rekruter
                  </h3>
                </div>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-3 gap-8 my-8 px-20">
                {/* Card 1 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md inline-block">
                      01
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Seleksi Kandidat Lebih Efektif
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Sistem akan membantu menyaring CV berdasarkan kriteria yang Anda tentukan.
                  </p>
                </div>

                {/* Card 2 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md inline-block">
                      02
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Proses Rekrutmen lebih Terstruktur
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Pantau pelamar dan kelola tahapan seleksi dengan lebih praktis.
                  </p>
                </div>

                {/* Card 3 */}
                <div className="bg-[#f9f5f0] px-6 py-8 rounded-lg shadow-md flex-col h-full">
                  <div className="flex items-start gap-3 mb-4">
                    <div className="bg-blue-900 text-white text-lg font-bold px-2 py-2 rounded-md">
                      03
                    </div>
                    <h3 className="text-xl font-semibold text-[#1e3a5f] mt-2">
                      Akses Kandidat yang Relevan
                    </h3>
                  </div>
                  <p className="text-black text-lg">
                    Temukan pelamar yang paling sesuai tanpa harus memilah satu per satu.
                  </p>
                </div>
              </div>
          </div>
            </div>
        </section>
      </main>
      <Footer />
    </div>
  )
}