"use client"

import React from "react"
import Image from "next/image"
import Link from "next/link"
import { FiSearch } from "react-icons/fi";
import { FaPlus } from "react-icons/fa6";
import { Button } from "@/components/ui/button";
import Footer from "@/components/global/footer";
import Navbar from "@/components/global/navbar";

type JobPosting = {
  id: number
  title: string
  company: string
  category: string
  location: string
  logo: string
  date: string
}

const jobList: JobPosting[] = [
  {
    id: 1,
    title: 'Backend Engineer',
    company: 'Shopee',
    category: 'Teknologi Informasi',
    location: 'Jakarta Pusat',
    logo: 'Shopeelogo.svg',
    date: '10 Maret 2025',
  },
  {
    id: 2,
    title: 'Frontend Engineer',
    company: 'Shopee',
    category: 'Teknologi Informasi',
    location: 'Jakarta Selatan',
    logo: 'Shopeelogo.svg',
    date: '15 Maret 2025',
  },
  {
    id: 3,
    title: 'UI/UX Designer',
    company: 'Shopee',
    category: 'Desain Produk',
    location: 'Jakarta Barat',
    logo: 'Shopeelogo.svg',
    date: '20 Maret 2025',
  },
  {
    id: 4,
    title: 'Product Manager',
    company: 'Shopee',
    category: 'Manajemen Produk',
    location: 'Jakarta Pusat',
    logo: 'Shopeelogo.svg',
    date: '25 Maret 2025',
  },
]

export default function Perekrut() {
    return (
      <div className="min-h-screen flex flex-col bg-white">
        {/* Header */}
        <Navbar role="perekrut"/>

        {/* Hero Section */}
        <div className="w-full bg-blue-100 py-25 text-center">
          <h2 className="text-3xl md:text-4xl font-bold text-[#1e3a5f]">Lihat, Kelola, dan Temukan Kandidat Terbaik!</h2>
          <p className="text-gray-700 mt-4 text-xl">
            Lihat, kelola, dan optimalkan proses rekrutmen Anda dengan lebih mudah.              
          </p>
        </div>

        {/* Search and Add Job */}
        <section className="py-6 bg-white">
          <div className="container mx-auto px-40">
            <div className="flex flex-col md:flex-row gap-4 justify-between">
              <form className="relative w-full max-w-5xl">
                <input
                  type="text"
                  placeholder="Cari Lowongan"
                  className="w-full px-4 py-2 border-2 border-gray-300 rounded-full focus:outline-none focus:ring focus:ring-[#1e3a5f]"
                />
                <button
                  type="submit"
                  className="absolute right-4 top-1/2 transform -translate-y-1/2 text-gray-500 hover:text-[#1e3a5f]"
                >
                  <FiSearch className="h-5 w-5" />
                </button>
              </form>

              <Link
                href="/perekrut/tambah-lowongan"
                className="flex items-center whitespace-nowrap justify-center gap-2 bg-[#1e3a5f] text-white px-5 py-2 rounded-full hover:bg-[#1e3a5f]/90 transition-colors"
              >
                <FaPlus className="h-5 w-5" />
                <span>tambah loker</span>
              </Link>
            </div>
          </div>
        </section>

        {/* Job Cards */}
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-x-6 gap-y-10 my-6 px-30">
          {jobList.map((job) => (
            <div key={job.id} className="border rounded-md p-4 w-full max-w-lg bg-white relative group border-black flex flex-col justify-between h-full">
              <Link href={`/job/${job.id}`} className="block flex-grow">
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
                <p className="font-semibold text-[#1e3a5f] text-lg mt-3">{job.category}</p>

                {/* Lokasi dan Tanggal */}
                <div className="text-black text-md">
                  <p>{job.location}</p>
                  <p>{job.date}</p>
                </div>
              </Link>
              
              {/* Tombol "Lihat Pelamar" muncul saat hover */}
              <div className="absolute left-0 right-0 -bottom-8 opacity-0 group-hover:opacity-100 transition-opacity duration-300 mt-auto">
                <Button className="w-full bg-[#f5efe7] text-gray-700 text-lg py-1 pt-2 text-left border border-black rounded-lg hover:bg-gray-100">
                lihat pelamar
                </Button>
              </div>
            </div>
          ))}
        </div>
          
        {/* Footer */}
        <Footer />
      </div>
    )
}