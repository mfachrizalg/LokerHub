"use client"

import type React from "react"
import { useRouter } from "next/navigation"
import { ArrowLeft, Plus, X } from "lucide-react"
import Image from "next/image"
import { addJobForm } from "@/hooks/tambahLowongan"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"

export default function AddJobPage() {
  const router = useRouter()

  const {
    formData,
    handleChange,
    newResponsibility,
    setNewResponsibility,
    addResponsibility,
    removeResponsibility,
    newQualification,
    setNewQualification,
    addQualification,
    removeQualification,
    handleSubmit,
    isSubmitting,
  } = addJobForm({
    title: "",
    jobType: "",
    positionLevel: "",
    salaryMin: "",
    salaryMax: "",
    field: "",
    closingDate: "",
    description: "",
    responsibilities: [""],
    qualifications: [""],
  })

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="perekrut"/>

      <div className="container mx-auto px-4 py-6 flex-1">
        {/* Back button */}
        <div className="mb-4">
          <button onClick={() => router.back()} className="flex items-center text-[#1e3a5f] hover:underline">
            <ArrowLeft className="h-4 w-4 mr-1" />
            Kembali
          </button>
        </div>

        <h1 className="text-2xl font-bold text-[#1e3a5f] text-center mb-6">Tambah Loker</h1>

        <form onSubmit={handleSubmit} className="max-w-4xl mx-auto">
          <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] p-6 mb-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {/* Kolom kiri */}
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Nama Pekerjaan</label>
                  <input
                    type="text"
                    name="title"
                    value={formData.title}
                    onChange={handleChange}
                    required
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Tipe Pekerjaan</label>
                  <input
                    type="text"
                    name="jobType"
                    value={formData.jobType}
                    onChange={handleChange}
                    required
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Tingkat Posisi</label>
                  <input
                    type="text"
                    name="positionLevel"
                    value={formData.positionLevel}
                    onChange={handleChange}
                    required
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Rentang Gaji</label>
                  <div className="flex items-center">
                    <input
                      type="text"
                      name="salaryMin"
                      value={formData.salaryMin}
                      onChange={handleChange}
                      required
                      className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                      placeholder="Min"
                    />
                    <span className="mx-2 text-gray-500">-</span>
                    <input
                      type="text"
                      name="salaryMax"
                      value={formData.salaryMax}
                      onChange={handleChange}
                      required
                      className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                      placeholder="Max"
                    />
                  </div>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Bidang</label>
                  <input
                    type="text"
                    name="field"
                    value={formData.field}
                    onChange={handleChange}
                    required
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Tanggal Penutupan</label>
                  <input
                    type="date"
                    name="closingDate"
                    value={formData.closingDate}
                    onChange={handleChange}
                    required
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  />
                </div>
              </div>

              {/* Kolom kanan */}
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi Pekerjaan</label>
                  <textarea
                    name="description"
                    value={formData.description}
                    onChange={handleChange}
                    required
                    rows={5}
                    className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                  ></textarea>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Tanggung Jawab</label>
                  <div className="flex items-center mb-2">
                    <input
                      type="text"
                      value={newResponsibility}
                      onChange={(e) => setNewResponsibility(e.target.value)}
                      className="flex-1 px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                      placeholder="Tambahkan tanggung jawab"
                    />
                    <button
                      type="button"
                      onClick={addResponsibility}
                      className="ml-2 bg-[#1e3a5f] text-white p-2 rounded-md hover:bg-[#1e3a5f]/90"
                    >
                      <Plus className="h-4 w-4" />
                    </button>
                  </div>
                  <div className="space-y-2 max-h-40 overflow-y-auto">
                    {formData.responsibilities.map(
                      (responsibility: string, index: number) =>
                        responsibility.trim() !== "" && (
                          <div key={index} className="flex items-center bg-white p-2 rounded-md border">
                            <span className="flex-1 text-sm">{responsibility}</span>
                            <button
                              type="button"
                              onClick={() => removeResponsibility(index)}
                              className="text-red-500 hover:text-red-700"
                            >
                              <X className="h-4 w-4" />
                            </button>
                          </div>
                        ),
                    )}
                  </div>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Kualifikasi</label>
                  <div className="flex items-center mb-2">
                    <input
                      type="text"
                      value={newQualification}
                      onChange={(e) => setNewQualification(e.target.value)}
                      className="flex-1 px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                      placeholder="Tambahkan kualifikasi"
                    />
                    <button
                      type="button"
                      onClick={addQualification}
                      className="ml-2 bg-[#1e3a5f] text-white p-2 rounded-md hover:bg-[#1e3a5f]/90"
                    >
                      <Plus className="h-4 w-4" />
                    </button>
                  </div>
                  <div className="space-y-2 max-h-40 overflow-y-auto">
                    {formData.qualifications.map(
                      (qualification: string, index: number) =>
                        qualification.trim() !== "" && (
                          <div key={index} className="flex items-center bg-white p-2 rounded-md border">
                            <span className="flex-1 text-sm">{qualification}</span>
                            <button
                              type="button"
                              onClick={() => removeQualification(index)}
                              className="text-red-500 hover:text-red-700"
                            >
                              <X className="h-4 w-4" />
                            </button>
                          </div>
                        ),
                    )}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="flex justify-center">
            <button
              type="submit"
              disabled={isSubmitting}
              className="bg-[#1e3a5f] text-white px-6 py-2 rounded-md hover:bg-[#1e3a5f]/90 transition-colors disabled:opacity-50"
            >
              {isSubmitting ? "Menyimpan..." : "simpan loker"}
            </button>
          </div>
        </form>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}