"use client"

import type React from "react"
import { useState, useEffect } from "react"
import Link from "next/link"
import { Search } from "lucide-react"
import Image from "next/image"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"

interface Job {
  id: string
  title: string
  company: string
  location: string
  date: string
  logo: string
  category: string
}

// mock data for job listings 
const mockJobs: Job[] = [
  {
    id: "job-1",
    title: "Backend Engineer",
    company: "Shopee",
    location: "Jakarta Pusat",
    date: "15 Maret 2025",
    logo: "/shopee-logo.svg",
    category: "Teknologi Informasi"
  },
  {
    id: "job-2",
    title: "Frontend Developer",
    company: "Tokopedia",
    location: "Jakarta Selatan",
    date: "10 Maret 2025",
    logo: "/tokopedia-logo.svg",
    category: "Teknologi Informasi"
  },
  {
    id: "job-3",
    title: "Data Scientist",
    company: "Gojek",
    location: "Bandung",
    date: "5 Maret 2025",
    logo: "/gojek-logo.svg",
    category: "Teknologi Informasi"
  },
  {
    id: "job-4",
    title: "Mobile Developer",
    company: "Bukalapak",
    location: "Yogyakarta",
    date: "8 Maret 2025",
    logo: "/bukalapak-logo.svg",
    category: "Teknologi Informasi"
  },
]

export default function Kandidat() {
  const [jobs, setJobs] = useState<Job[]>(mockJobs)
  const [searchQuery, setSearchQuery] = useState("")
  const [currentPage, setCurrentPage] = useState(1)
  const [isLoading, setIsLoading] = useState(false)

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    setCurrentPage(1) // Reset to first page on new search
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="kandidat"/>

      {/* Hero section with search */}
      <section className="bg-[#a8c5e2] py-25 md:py-18">
        <div className="container mx-auto px-4">
          <h1 className="text-4xl font-bold text-[#1e3a5f] text-center mb-2">Pekerjaan yang Cocok Untukmu</h1>
          <p className="text-center text-xl text-[#1e3a5f] mb-6">
            Jangan Biarkan Kesempatan Berlalu, Dapatkan Karier Terbaik untuk Masa Depanmu!
          </p>
          <div className="px-20 md:px-4 pt-6">
            <form onSubmit={handleSearch} className="max-w-2xl mx-auto relative">
              <input
                type="text"
                placeholder="Cari Pekerjaan"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full px-4 py-2 rounded-full border-2 bg-[#f5fef7] border-black focus:outline-none"
              />
              <button
                type="submit"
                className="absolute right-5 top-1/2 transform -translate-y-1/2 text-gray-500 hover:text-[#1e3a5f]"
              >
                <Search className="h-5 w-5" />
              </button>
            </form>
          </div>          
        </div>
      </section>

      {/* Job listings */}
      <section className="py-8">
        <div className="container mx-auto px-4">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-x-6 gap-y-12">
            {jobs.map((job) => (
              <div key={job.id} className="relative group">
                <div className="bg-white rounded-lg shadow-md p-4 border border-black hover:shadow-lg transition-shadow relative z-20">
                  <Link href={`/job/${job.id}`} className="block">
                    {/* Logo + Company Info */}
                    <div className="flex items-center">
                      <div className="bg-red-100 rounded-md p-2 mr-3">
                        <Image
                          src={job.logo}
                          alt={job.company}
                          width={40}
                          height={40}
                          className="rounded-md"
                        />
                      </div>
                      <div>
                        <h3 className="font-bold text-xl text-[#1e3a5f]">{job.title}</h3>
                        <p className="text-xl text-gray-600">{job.company}</p>
                      </div>
                    </div>

                    {/* Bidang */}
                    <p className="font-bold text-[#1e3a5f] text-lg mt-3">{job.category}</p>

                    {/* Lokasi dan Tanggal */}
                    <div className="text-black text-lg">
                      <p>{job.location}</p>
                      <p>{job.date}</p>
                    </div>
                  </Link>
                </div>

                {/* Hover: Jumlah Pelamar */}
                <div className="absolute left-0 right-0 -bottom-8 opacity-0 group-hover:opacity-100 transition-opacity duration-300 z-10">
                  <div className="bg-[#f5efe7] text-gray-700 text-lg py-1 pt-2 text-center border border-black rounded-lg">
                    <p>2 pelamar sudah mendaftar di loker ini</p>
                  </div>
                </div>
              </div>
            ))}
          </div>

          {/* Pagination */} 
          <div className="flex justify-center mt-8">
            <nav className="flex space-x-2">
              {[1, 2, 3].map((page) => (
                <button
                  key={page}
                  onClick={() => setCurrentPage(page)}
                  className={`px-3 py-1 rounded ${
                    currentPage === page ? "bg-[#1e3a5f] text-white" : "bg-gray-200 text-gray-700 hover:bg-gray-300"
                  }`}
                >
                  {page}
                </button>
              ))}
            </nav>
          </div>
        </div>
      </section>

      {/* Footer */}
      <Footer />
    </div>
  )
}
