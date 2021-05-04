package sexpr

import (
	"bufio"
	"encoding/binary"
	"io"
)

const (
	SAMPLERATE       = 8e3 // 8kHz audio
	CHANNELS         = 2   // Stereo
	BYTES_PER_SECOND = CHANNELS * SAMPLERATE
)

func WriteWav(f Fn, w io.Writer, duration int) error {
	buf := bufio.NewWriter(w)
	nsamples := duration * BYTES_PER_SECOND

	err := writeWaveFileHeader(buf, waveFileHeaderOptions{
		DataSize:    nsamples,
		NumChannels: CHANNELS,
		SampleRate:  SAMPLERATE,
		Precision:   1,
	})
	if err != nil {
		return err
	}

	// Write actual bytes of audio
	for i := 0; i < nsamples; i++ {
		err := buf.WriteByte(byte(f.Eval(i)))
		if err != nil {
			return err
		}
	}
	err = buf.Flush()
	if err != nil {
		return err
	}

	return nil
}

type waveFileHeader struct {
	RiffMark      [4]byte
	FileSize      int32
	WaveMark      [4]byte
	FmtMark       [4]byte
	FormatSize    int32
	FormatType    int16
	NumChans      int16
	SampleRate    int32
	ByteRate      int32
	BytesPerFrame int16
	BitsPerSample int16
	DataMark      [4]byte
	DataSize      int32
}

type waveFileHeaderOptions struct {
	DataSize, NumChannels, SampleRate, Precision int
}

func writeWaveFileHeader(w io.Writer, opts waveFileHeaderOptions) error {
	h := waveFileHeader{
		RiffMark:      [4]byte{'R', 'I', 'F', 'F'},
		FileSize:      int32(opts.DataSize + 44), /* 44 is the header size */
		WaveMark:      [4]byte{'W', 'A', 'V', 'E'},
		FmtMark:       [4]byte{'f', 'm', 't', ' '},
		FormatSize:    16,
		FormatType:    1,
		NumChans:      int16(opts.NumChannels),
		SampleRate:    int32(opts.SampleRate),
		ByteRate:      int32(int(opts.SampleRate) * opts.NumChannels * opts.Precision),
		BytesPerFrame: int16(opts.NumChannels * opts.Precision),
		BitsPerSample: int16(opts.Precision) * 8,
		DataMark:      [4]byte{'d', 'a', 't', 'a'},
		DataSize:      int32(opts.DataSize),
	}
	if err := binary.Write(w, binary.LittleEndian, &h); err != nil {
		return err
	}
	return nil
}
