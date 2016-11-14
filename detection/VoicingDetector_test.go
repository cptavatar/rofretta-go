package detection

import "testing"

func TestVoicingDetector_ParseAndIdentifyVoicing(t *testing.T) {
	vd := CreateVoicingDetector()
	voicing, valid := vd.ParseAndIdentifyVoicing("032010")
	if !valid {
		t.Error("ParseAndIdentify was passed valid string but returned invalid")
	}
	if len(voicing) != 1 {
		t.Errorf("We should have 1 voicing identified, recieved %d", len(voicing))
	}
}
