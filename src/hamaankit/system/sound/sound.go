package sound

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type SoundF []byte
type Player *audio.Player

func SetVolPlayer(v float64, m *Player) {
	var mm *audio.Player
	mm = *m
	mm.SetVolume(v)
}

const (
	SAMPLE_RATE = 32000 //44100
)

var audioContext = audio.NewContext(SAMPLE_RATE)

func PlaySE(b SoundF, format string) {
	// PlayBGM(b, format)
	switch format {
	case "mp3":

		s, err := mp3.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(b))
		if err != nil {
			log.Fatal(err)
		}

		p, err := audioContext.NewPlayer(s)

		if err != nil {
			log.Fatal(err)
		}

		p.Play()
	case "ogg":
		s, err := vorbis.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(b))
		if err != nil {
			log.Fatal(err)
		}

		p, err := audioContext.NewPlayer(s)

		if err != nil {
			log.Fatal(err)
		}

		p.Play()
	case "wav":

		s, err := wav.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(b))
		// fmt.Println("ぬ")
		if err != nil {
			log.Fatal(err)
		}

		p, err := audioContext.NewPlayer(s)

		if err != nil {
			log.Fatal(err)
		}

		p.Play()
	}
}

func PlayBGM(bM SoundF, format string) (kuro Player) {
	if format == "wav" {
		return PlayBGMWAV(bM)
	} else if format == "mp3" {

		sM, err := mp3.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(bM))
		if err != nil {
			log.Fatal(err)
		}

		pM, err := audioContext.NewPlayer(sM)

		if err != nil {
			log.Fatal(err)
		}

		pM.Play()
		//for{

		//}
		return pM
	} else if format == "ogg" {

		sM, err := vorbis.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(bM))
		if err != nil {
			log.Fatal(err)
		}

		pM, err := audioContext.NewPlayer(sM)

		if err != nil {
			log.Fatal(err)
		}

		pM.Play()
		return pM
	} else {
		return
	}

}

func PlayBGMWAV(bM SoundF) (kuro Player) {

	sM, err := wav.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(bM))
	if err != nil {
		log.Fatal(err)
	}

	pM, err := audioContext.NewPlayer(sM)

	if err != nil {
		log.Fatal(err)
	}

	pM.Play()
	//for{

	//}
	return pM

}

func Close(get *audio.Player) error {
	return get.Close()
}

func Check(get2 *audio.Player) bool {
	return get2.IsPlaying()
}

var BattEm embed.FS

// 音声読み込み
func LoadSound(SoundPointer *SoundF, Path string) {
	var SndLoaderr error

	// fmt.Println(BattEm)
	*SoundPointer, SndLoaderr = BattEm.ReadFile(Path)
	if SndLoaderr != nil {
		log.Fatal(SndLoaderr)
	}
}
